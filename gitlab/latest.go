package gitlab

import (
	"fmt"
	"regexp"
	"sort"

	"github.com/MamaShip/MR-Tracker/utils"
	"github.com/blang/semver/v4"
)

// Given a tag, return a (start, end) tag tuple
//   end = the given tag
//   start = latest formal version tag before the given tag
func GetLatestTag(s utils.UserSettings) (string, string) {
	var g Gitlab
	if isOfficialGitlab(s.Site) {
		g = NewGitlab(s.Project, s.Token)
	} else {
		g = NewCustomGitlab(s.Site, s.Project, s.Token)
	}

	tags := g.getTags()

	sort.Slice(tags, func(i, j int) bool { // tag 升序排列
		itime, _ := utils.ParseTime(tags[i].Commit.CreatedAt)
		jtime, _ := utils.ParseTime(tags[j].Commit.CreatedAt)
		return itime.Before(jtime)
	})

	var start, end string
	if tagExist(s.Latest, tags) {
		end = s.Latest
		start = prevSemver(end, tags)
	}
	return start, end
}

func prevSemver(current string, tags []Tag) string {
	endSemver, err := getSemver(current)
	if err != nil {
		return ""
	}
	// use semver.Version struct to compare versions
	latestSemver, _ := semver.Make("0.0.0") // fake initial version
	var latestTag string
	for _, tag := range tags {
		if tag.Name == current {
			continue
		}
		semver, err := getSemver(tag.Name)
		if err != nil {
			continue
		}
		if semver.GT(latestSemver) && semver.LT(endSemver) {
			if isFormalVersion(tag.Name) {
				latestSemver = semver
				latestTag = tag.Name
			}
		}
	}
	return latestTag
}

func tagExist(want string, tags []Tag) bool {
	for _, tag := range tags {
		if want == tag.Name {
			return true
		}
	}
	return false
}

var semverPattern = regexp.MustCompile(`[a-zA-Z]?(\d+\.\d+\.\d+).*`)

// extract "0.0.0" part from given string
// and convert it to semver.Version struct
func getSemver(s string) (semver.Version, error) {
	matched := semverPattern.FindStringSubmatch(s)
	if (matched != nil) && (len(matched) >= 2) {
		return semver.Parse(matched[1])
	}
	return semver.Version{}, fmt.Errorf("invalid semver")
}

var formalVerRgx = regexp.MustCompile(`^[a-zA-Z]?(\d+\.\d+\.\d+)$`)

func isFormalVersion(s string) bool {
	return formalVerRgx.MatchString(s)
}

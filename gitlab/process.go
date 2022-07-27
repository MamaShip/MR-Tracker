package gitlab

import (
	"fmt"
	"strings"

	"github.com/MamaShip/MR-Tracker/utils"
)

func isOfficialGitlab(url string) bool {
	return strings.Contains(url, "gitlab.com")
}

func FetchMrs(s utils.UserSettings) ([]MergeRequest, error) {
	var g Gitlab
	if isOfficialGitlab(s.Site) {
		g = NewGitlab(s.Project, s.Token)
	} else {
		g = NewCustomGitlab(s.Site, s.Project, s.Token)
	}

	if s.Branch == "" {
		if branch, err := g.getDefaultBranch(); err != nil {
			return nil, err
		} else {
			s.Branch = branch.Name
		}
	}

	return g.FindMRsBetween(s.StartTag, s.EndTag, s.Branch)
}

func Post2Issue(changes string, s utils.UserSettings) error {
	if changes == "" {
		fmt.Println("Nothing to post")
		return nil
	}

	title := utils.GenerateTitle(s.StartTag, s.EndTag)

	var g Gitlab
	if isOfficialGitlab(s.Site) {
		g = NewGitlab(s.Project, s.Token)
	} else {
		g = NewCustomGitlab(s.Site, s.Project, s.Token)
	}

	if g.issueExists(title) {
		return g.updateIssue(title, changes)
	} else {
		return g.newIssue(title, changes)
	}
}

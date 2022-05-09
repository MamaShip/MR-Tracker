package gitlab

import (
	"encoding/json"
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

func findTag(tags []Tag, tag_name string) (Tag, error) {
	for _, tag := range tags {
		if tag.Name == tag_name {
			return tag, nil
		}
	}
	return Tag{}, fmt.Errorf("tag %s not found", tag_name)
}

func Post2Issue(changes string, s utils.UserSettings) error {
	var g Gitlab
	if isOfficialGitlab(s.Site) {
		g = NewGitlab(s.Project, s.Token)
	} else {
		g = NewCustomGitlab(s.Site, s.Project, s.Token)
	}

	issue_api := g.String() + "/issues"
	r := IssueRqst{
		Title: fmt.Sprintf("[MR Tracker] Changes between %s - %s",
			s.StartTag,
			s.EndTag),
		Description: changes,
		Token:       g.Token,
	}
	jsonData, err := json.Marshal(r)
	if err != nil {
		fmt.Println(err)
		return err
	}
	response := utils.Post(issue_api, jsonData)
	result := ParseIssueResp(response)
	fmt.Printf("Issue %d created: %s\n", result.Id, result.WebUrl)
	return nil
}

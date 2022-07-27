package gitlab

import (
	"encoding/json"
	"fmt"
	"net/url"
	"time"

	"github.com/MamaShip/MR-Tracker/utils"
)

func (g *Gitlab) getIssues() []Issue {
	listIssueApi := g.String() + "/issues"
	p := url.Values{}
	p.Set("private_token", g.Token)
	p.Set("per_page", "100")
	issues := make([]Issue, 0)
	total_page := 1
	for i := 1; i <= total_page; i++ {
		p.Set("page", fmt.Sprintf("%d", i))
		getIssue := utils.FormRequest(listIssueApi, p)
		var json_str []byte
		json_str, total_page = utils.Get(getIssue)
		paged_issues := ParseIssueList(json_str)
		issues = append(issues, paged_issues...)
	}
	return issues
}

func (g *Gitlab) issueExists(title string) bool {
	issueList := g.getIssues()
	for _, issue := range issueList {
		if issue.Title == title {
			return true
		}
	}
	return false
}

func (g *Gitlab) newIssue(title string, body string) error {
	r := IssueRqst{
		Title:       title,
		Description: body,
		Token:       g.Token,
	}
	jsonData, err := json.Marshal(r)
	if err != nil {
		fmt.Println(err)
		return err
	}
	issue_api := g.String() + "/issues"
	response := utils.Post(issue_api, jsonData)
	result := ParseIssue(response)
	fmt.Printf("Issue %d created: %s\n", result.Id, result.WebUrl)
	return nil
}

func (g *Gitlab) updateIssue(title string, body string) error {
	find := false
	var old Issue
	issueList := g.getIssues()
	for _, issue := range issueList {
		if issue.Title == title {
			find = true
			old = issue
			break
		}
	}
	if !find {
		return fmt.Errorf("can't find issue")
	}

	r := IssueRqst{
		Title:       title,
		Description: fmt.Sprintf("Re-generated at: %s\n\n%s", time.Now().Format("2006-01-02 15:04:05"), body),
		Token:       g.Token,
		StateEvent:  "reopen",
	}
	jsonData, err := json.Marshal(r)
	if err != nil {
		fmt.Println(err)
		return err
	}

	issue_api := fmt.Sprintf("%s/issues/%d", g.String(), old.Iid)
	response := utils.Put(issue_api, jsonData)
	result := ParseIssue(response)
	fmt.Printf("Issue %d updated: %s\n", result.Id, result.WebUrl)
	return nil
}

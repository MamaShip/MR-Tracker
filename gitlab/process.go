package gitlab

import (
	"encoding/json"
	"fmt"
	"net/url"
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

	return g.FindMRsBetween(s.StartTag, s.EndTag)
}

func (g *Gitlab) FindMRsBetween(start_tag string, end_tag string) ([]MergeRequest, error) {
	tags := g.getTags()

	start, err := findTag(tags, start_tag)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	end, err := findTag(tags, end_tag)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	// fmt.Println(start.Commit.CreatedAt, end.Commit.CreatedAt)
	mrs := g.getMRsAfter(start.Commit.CreatedAt)
	mrs = filterMRs(mrs, start, end)

	return mrs, nil
}

func (g *Gitlab) getTags() []Tag {
	tag_api := g.String() + "/repository/tags"
	p := url.Values{}
	p.Set("private_token", g.Token)
	get_tag := utils.FormRequest(tag_api, p)
	// fmt.Println(get_tag)
	json_str := utils.Get(get_tag)
	return ParseTags(json_str)
}

func findTag(tags []Tag, tag_name string) (Tag, error) {
	for _, tag := range tags {
		if tag.Name == tag_name {
			return tag, nil
		}
	}
	return Tag{}, fmt.Errorf("tag %s not found", tag_name)
}

func (g *Gitlab) getMRsAfter(start_time string) []MergeRequest {
	mr_api := g.String() + "/merge_requests"
	p := url.Values{}
	p.Set("private_token", g.Token)
	p.Set("state", "merged")
	p.Set("order_by", "updated_at")
	p.Set("sort", "desc")
	p.Set("scope", "all")
	p.Set("target_branch", "master")
	p.Set("updated_after", start_time)
	get_mr := utils.FormRequest(mr_api, p)
	// fmt.Println(get_mr)
	json_str := utils.Get(get_mr)
	return ParseMRs(json_str)
}

// 根据始末 tag 过滤 merge request。
// 过滤结果剔除了 start tag 指向的 MR。包含了 end tag 指向的 MR。
func filterMRs(all_mr []MergeRequest, start Tag, end Tag) []MergeRequest {
	start_time, err := utils.ParseTime(start.Commit.CreatedAt)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	end_time, err := utils.ParseTime(end.Commit.CreatedAt)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	mrs := make([]MergeRequest, 0, len(all_mr))
	for _, mr := range all_mr {
		merge_time, err := utils.ParseTime(mr.MergedAt)
		if err != nil {
			fmt.Println(err)
			return nil
		}
		if merge_time.After(start_time) && merge_time.Before(end_time) {
			if mr.MergeCommit != start.Commit.Id {
				mrs = append(mrs, mr)
			}
		} else if mr.MergeCommit == end.Commit.Id {
			mrs = append(mrs, mr)
		}
	}
	return mrs
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

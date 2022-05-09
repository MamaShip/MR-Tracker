package gitlab

import (
	"fmt"
	"net/url"

	"github.com/MamaShip/MR-Tracker/utils"
)

type Gitlab struct {
	url.URL
	Token string
}

func NewGitlab(project_id int, token string) Gitlab {
	instance := Gitlab{}
	instance.Scheme = "http"
	instance.Host = "gitlab.com"
	instance.Path = fmt.Sprintf("/api/v4/projects/%d", project_id)
	instance.Token = token
	return instance
}

func NewCustomGitlab(host string, project_id int, token string) Gitlab {
	instance := Gitlab{}
	instance.Scheme = "http" // default to http, some gitlab instance not working on TLS
	instance.Host = host
	instance.Path = fmt.Sprintf("/api/v4/projects/%d", project_id)
	instance.Token = token
	return instance
}

func (g *Gitlab) FindMRsBetween(start_tag string, end_tag string, br string) ([]MergeRequest, error) {
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

	mrs := g.getMRsAfter(start.Commit.CreatedAt, br)
	mrs = filterMRs(mrs, start, end)
	return mrs, nil
}

func (g *Gitlab) getTags() []Tag {
	tag_api := g.String() + "/repository/tags"
	p := url.Values{}
	p.Set("private_token", g.Token)
	get_tag := utils.FormRequest(tag_api, p)
	json_str := utils.Get(get_tag)
	return ParseTags(json_str)
}

func (g *Gitlab) getMRsAfter(start_time string, br string) []MergeRequest {
	mr_api := g.String() + "/merge_requests"
	p := url.Values{}
	p.Set("private_token", g.Token)
	p.Set("state", "merged")
	p.Set("order_by", "updated_at")
	p.Set("sort", "desc")
	p.Set("scope", "all")
	p.Set("target_branch", br)
	p.Set("updated_after", start_time)
	get_mr := utils.FormRequest(mr_api, p)
	json_str := utils.Get(get_mr)
	return ParseMRs(json_str)
}

func (g *Gitlab) getBranches() []Branch {
	mr_api := g.String() + "/repository/branches"
	p := url.Values{}
	p.Set("private_token", g.Token) // token is not necessary for public repo
	get_mr := utils.FormRequest(mr_api, p)
	json_str := utils.Get(get_mr)
	return ParseBranches(json_str)
}

func (g *Gitlab) getDefaultBranch() (Branch, error) {
	branches := g.getBranches()
	for _, b := range branches {
		if b.Default {
			return b, nil
		}
	}
	return Branch{}, fmt.Errorf("no default branch found")
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

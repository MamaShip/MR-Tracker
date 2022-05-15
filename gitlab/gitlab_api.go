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

func (g *Gitlab) getTags() []Tag {
	tag_api := g.String() + "/repository/tags"
	p := url.Values{}
	p.Set("private_token", g.Token)
	p.Set("per_page", "100")
	tags := make([]Tag, 0)
	total_page := 1
	for i := 1; i <= total_page; i++ {
		p.Set("page", fmt.Sprintf("%d", i))
		get_tag := utils.FormRequest(tag_api, p)
		var json_str []byte
		json_str, total_page = utils.Get(get_tag)
		paged_tags := ParseTags(json_str)
		tags = append(tags, paged_tags...)
	}
	return tags
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
	p.Set("per_page", "100")
	mrs := make([]MergeRequest, 0)
	total_page := 1
	for i := 1; i <= total_page; i++ {
		p.Set("page", fmt.Sprintf("%d", i))
		get_mr := utils.FormRequest(mr_api, p)
		var json_str []byte
		json_str, total_page = utils.Get(get_mr)
		paged_mrs := ParseMRs(json_str)
		mrs = append(mrs, paged_mrs...)
	}
	return mrs
}

func (g *Gitlab) getAllMRs(br string) []MergeRequest {
	mr_api := g.String() + "/merge_requests"
	p := url.Values{}
	p.Set("private_token", g.Token)
	p.Set("state", "merged")
	p.Set("order_by", "updated_at")
	p.Set("sort", "desc")
	p.Set("scope", "all")
	p.Set("target_branch", br)
	p.Set("per_page", "100")
	mrs := make([]MergeRequest, 0)
	total_page := 1
	for i := 1; i <= total_page; i++ {
		p.Set("page", fmt.Sprintf("%d", i))
		get_mr := utils.FormRequest(mr_api, p)
		var json_str []byte
		json_str, total_page = utils.Get(get_mr)
		paged_mrs := ParseMRs(json_str)
		mrs = append(mrs, paged_mrs...)
	}
	return mrs
}

func (g *Gitlab) getBranches() []Branch {
	br_api := g.String() + "/repository/branches"
	p := url.Values{}
	p.Set("private_token", g.Token) // token is not necessary for public repo
	p.Set("per_page", "100")
	brs := make([]Branch, 0)
	total_page := 1
	for i := 1; i <= total_page; i++ {
		p.Set("page", fmt.Sprintf("%d", i))
		get_br := utils.FormRequest(br_api, p)
		var json_str []byte
		json_str, total_page = utils.Get(get_br)
		paged_brs := ParseBranches(json_str)
		brs = append(brs, paged_brs...)
	}
	return brs
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

func (g *Gitlab) findTag(tag_name string) (Tag, error) {
	tags := g.getTags()
	for _, tag := range tags {
		if tag.Name == tag_name {
			return tag, nil
		}
	}
	return Tag{}, fmt.Errorf("tag %s not found", tag_name)
}

func (g *Gitlab) getStartEndTag(start string, end string) (Tag, Tag, error) {
	var start_tag, end_tag Tag
	var err error
	if start != "" {
		start_tag, err = g.findTag(start)
		if err != nil {
			fmt.Println(err)
			return start_tag, end_tag, err
		}
	}
	if end != "" {
		end_tag, err = g.findTag(end)
		if err != nil {
			fmt.Println(err)
			return start_tag, end_tag, err
		}
	}
	return start_tag, end_tag, nil
}

func (g *Gitlab) FindMRsBetween(start string, end string, br string) ([]MergeRequest, error) {
	start_tag, end_tag, err := g.getStartEndTag(start, end)
	if err != nil {
		return nil, err
	}

	var mrs []MergeRequest
	if start_tag.Name != "" {
		mrs = g.getMRsAfter(start_tag.Commit.CreatedAt, br)
	} else {
		mrs = g.getAllMRs(br)
	}
	mrs = keepMRsBetween(mrs, start_tag, end_tag)
	return mrs, nil
}

package gitlab

import (
	"net/url"
	"testing"

	"github.com/MamaShip/MR-Tracker/utils"
)

func TestForm(t *testing.T) {
	u := "https://gitlab.example.com/api/v4/projects"
	params := url.Values{}
	params.Set("state", "merged")
	params.Set("order_by", "updated_at")
	params.Set("sort", "desc")
	params.Set("scope", "all")
	params.Set("target_branch", "master")
	result := FormRequest(u, params)
	println(result)
}

func TestNewGitlab(t *testing.T) {
	g := NewGitlab()
	println(g.String())
}

func TestAPI(t *testing.T) {
	g := NewCustomGitlab("gitlab.qitantech.com", "")
	// println(g.String())
	API := g.String() + "/merge_requests"
	p := url.Values{}
	p.Set("state", "merged")
	p.Set("scope", "all")
	// p.Set("private_token", "xxxxx")
	req := FormRequest(API, p)
	println(req)
	println(utils.Get(req))
}
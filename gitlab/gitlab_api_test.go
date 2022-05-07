package gitlab

import (
	"net/url"
	"testing"

	"github.com/MamaShip/MR-Tracker/utils"
)

func TestNewGitlab(t *testing.T) {
	g := NewGitlab(0, "")
	println(g.String())
}

func TestAPI(t *testing.T) {
	g := NewCustomGitlab("gitlab.qitantech.com", 102, "VhnrgrMbb51t9P3c3ZtG")
	API := g.String() + "/merge_requests"
	p := url.Values{}
	p.Set("state", "merged")
	p.Set("scope", "all")

	req := utils.FormRequest(API, p)
	println(req)
	println(string(utils.Get(req)))
}
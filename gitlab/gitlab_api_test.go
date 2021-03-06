package gitlab

import (
	"net/url"
	"testing"

	"github.com/MamaShip/MR-Tracker/utils"
	"github.com/stretchr/testify/assert"
)

func TestNewGitlab(t *testing.T) {
	g := NewGitlab(0, "")
	println(g.String())
}

func TestAPI(t *testing.T) {
	g := NewCustomGitlab("gitlab.com", 31285645, "")
	API := g.String() + "/merge_requests"
	p := url.Values{}
	p.Set("state", "merged")
	p.Set("scope", "all")

	req := utils.FormRequest(API, p)
	println(req)
	resp, _ := utils.Get(req)
	println(string(resp))
}

func TestDefaultBranch(t *testing.T) {
	g := NewCustomGitlab("gitlab.com", 31285645, "")
	br, err := g.getDefaultBranch()
	assert.NoError(t, err)
	assert.Equal(t, "master", br.Name)
}

func TestGetAllMRs(t *testing.T) {
	g := NewCustomGitlab("gitlab.com", 31285645, "")
	mrs := g.getAllMRs("master")
	for _, mr := range mrs {
		println(mr.Title)
	}
	println(len(mrs))
}

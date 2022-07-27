package gitlab

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_getIssues(t *testing.T) {
	g := NewGitlab(31285645, "")
	issues := g.getIssues()
	for _, i := range issues {
		println(i.Title, i.State)
	}
	println(len(issues))
	assert.Greater(t, len(issues), 0)
}

func Test_issueExists(t *testing.T) {
	g := NewGitlab(31285645, "")
	exists := g.issueExists("Test Issue")
	assert.Equal(t, true, exists)
	exists = g.issueExists("Nothing")
	assert.Equal(t, false, exists)
}

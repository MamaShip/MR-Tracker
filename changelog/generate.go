package changelog

import (
	"fmt"
	"strings"

	"github.com/MamaShip/MR-Tracker/gitlab"
)

func GenerateChanglog(mrs []gitlab.MergeRequest) string{
	lines := make([]string, 0)
	lines = append(lines, "## Changes\n")
	for _, mr := range mrs {
		line := fmt.Sprintf("- [!%d](%s) %s by %s", mr.IId, mr.MergeUrl, mr.Title, mr.Author.Name)
		lines = append(lines, line)
	}
	output := strings.Join(lines, "\n")
	fmt.Println(output)
	return output
}

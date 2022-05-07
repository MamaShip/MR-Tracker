package changelog

import (
	"fmt"
	"strings"

	"github.com/MamaShip/MR-Tracker/gitlab"
)

func GenerateChanglog(mrs []gitlab.MergeRequest) string{
	lines := make([]string, 0)
	for _, mr := range mrs {
		line := fmt.Sprintf("- [!%d](%s) %s by %s", mr.IId, mr.MergeUrl, mr.Title, mr.Author.Name)
		lines = append(lines, line)
	}
	output := strings.Join(lines, "\n")
	fmt.Println(">> Found changes:")
	fmt.Println(output)
	return output
}

package changelog

import (
	"fmt"
	"strings"

	"github.com/MamaShip/MR-Tracker/gitlab"
	"github.com/MamaShip/MR-Tracker/utils"
)

func GenerateFullVer(mrs []gitlab.MergeRequest) string {
	if len(mrs) == 0 {
		return ""
	}
	fmt.Println(">> Found changes:")
	lines := make([]string, 0)
	for _, mr := range mrs {
		fmt.Printf("- [!%d] %s by %s\n", mr.IId, mr.Title, mr.Author.Name)
		merge_time, _ := utils.ParseTime(mr.MergedAt)
		line := fmt.Sprintf("- [!%d](%s) %s -- [%s by [%s](%s)]",
			mr.IId,
			mr.MergeUrl,
			mr.Title,
			merge_time.Format("01-02"),
			mr.Author.Name,
			mr.Author.Url)
		lines = append(lines, line)
	}
	output := strings.Join(lines, "\n")
	return output
}

func GenerateSimpleVer(mrs []gitlab.MergeRequest) string {
	if len(mrs) == 0 {
		return ""
	}
	fmt.Println(">> Found changes:")
	lines := make([]string, 0)
	for _, mr := range mrs {
		fmt.Printf("- [!%d] %s\n", mr.IId, mr.Title)
		line := fmt.Sprintf("- [!%d](%s) %s",
			mr.IId,
			mr.MergeUrl,
			mr.Title)
		lines = append(lines, line)
	}
	output := strings.Join(lines, "\n")
	return output
}

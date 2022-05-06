package changelog

import (
	"fmt"

	"github.com/MamaShip/MR-Tracker/gitlab"
)

func GenerateChanglog(mrs []gitlab.MergeRequest) {
	fmt.Println("## Changes")
	for _, mr := range mrs {
		fmt.Println("- ", mr.Title)
	}
}

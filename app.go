package main

import (
	"github.com/MamaShip/MR-Tracker/changelog"
	"github.com/MamaShip/MR-Tracker/gitlab"
)

func main() {
	println("Start Tracking!")
	mrs, err := gitlab.FindMRsBetween("v1.9.3", "v1.9.4")
	if err != nil {
		println(err)
		return
	}
	changelog.GenerateChanglog(mrs)
}

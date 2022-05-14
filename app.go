package main

import (
	"flag"
	"fmt"

	"github.com/MamaShip/MR-Tracker/changelog"
	"github.com/MamaShip/MR-Tracker/gitlab"
	"github.com/MamaShip/MR-Tracker/utils"
)

func main() {
	// parse command line arguments
	flag.Parse()
	// if `-v`, print version and exit
	utils.PrintVersion()
	if utils.RequestVersion {
		return
	}

	// avoid invalid settings
	err := utils.CheckSettings()
	if err != nil {
		fmt.Println(err)
		return
	}

	mrs, err := gitlab.FetchMrs(utils.Settings)
	if err != nil {
		fmt.Println(err)
		return
	}

	var changes string
	if utils.Settings.Simplify {
		changes = changelog.GenerateSimpleVer(mrs)
	} else {
		changes = changelog.GenerateFullVer(mrs)
	}

	if utils.Settings.PostIssue {
		println(">> Posting changes to issue...")
		if err := gitlab.Post2Issue(changes, utils.Settings); err != nil {
			fmt.Println(err)
			return
		} else {
			println(">> Post issue successfully!")
		}
	}

	if utils.Settings.Output != "" {
		err := utils.Write2File(changes, utils.Settings.Output)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

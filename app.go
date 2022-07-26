package main

import (
	"flag"
	"fmt"
	"log"

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
		log.Fatal(err)
	}

	if utils.Settings.Latest != "" {
		start, end := gitlab.GetLatestTag(utils.Settings)
		if start == "" && end == "" {
			log.Fatal("'-latest' enabled but no valid tag found!")
		}
		utils.Settings.StartTag, utils.Settings.EndTag = start, end
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
			fmt.Println(err, " Fail to post issue")
		} else {
			println(">> Post issue successfully!")
		}
	}

	if utils.Settings.Output != "" {
		output2File(changes, utils.Settings)
	}
}

func output2File(changes string, s utils.UserSettings) {
	title := utils.GenerateTitle(s.StartTag, s.EndTag)
	text := "# " + title + "\n\n" + changes
	err := utils.Write2File(text, s.Output)
	if err != nil {
		fmt.Println(err, " Fail to write file")
	}
}

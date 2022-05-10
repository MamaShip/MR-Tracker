package main

import (
	"flag"
	"fmt"

	"github.com/MamaShip/MR-Tracker/changelog"
	"github.com/MamaShip/MR-Tracker/gitlab"
	"github.com/MamaShip/MR-Tracker/utils"
)

func main() {
	// 解析命令行参数
	flag.Parse()
	utils.PrintVersion()
	// 如果用户在查询版本号，显示完后就直接退出
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
	changes := changelog.GenerateChanglog(mrs)

	if utils.Settings.PostIssue {
		println(">> Posting changes to issue...")
		if err := gitlab.Post2Issue(changes, utils.Settings); err != nil {
			fmt.Println(err)
			return
		} else {
			println(">> Post issue successfully!")
		}
	}
}

package utils

import (
	"flag"
	"fmt"
	"os"
)

// 用户设定的 run time 参数
type UserSettings struct {
	// Port                int
	Site      string
	Project   int
	StartTag  string
	EndTag    string
	PostIssue bool
	Token     string
}

var (
	UserRequestVerion bool
	Settings          UserSettings
)

func init() {
	// 注意相同的变量 UserRequestVerion 绑定了不同的参数，这是允许的
	// 但要确保它们的默认值完全一致，否则会有预期之外的错误
	flag.BoolVar(&UserRequestVerion, "version", false, "与 -v 相同")
	flag.BoolVar(&UserRequestVerion, "v", false, "显示版本号")

	// flag.IntVar(&Port, "p", 50051, "监听的端口号")
	flag.StringVar(&Settings.StartTag, "start", "", "作为统计起点的tag")
	flag.StringVar(&Settings.EndTag, "end", "", "作为统计终点的tag")
	flag.BoolVar(&Settings.PostIssue, "post", false, "把版本差异记录到 gitlab issue")

	flag.StringVar(&Settings.Site, "site", "gitlab.com", "gitlab 实例的域名（default: gitlab.com）")
	flag.IntVar(&Settings.Project, "project", 0, "Project ID")
	flag.StringVar(&Settings.Token, "token", "", "拥有对应仓库权限的 token(see: https://docs.gitlab.com/ee/user/project/settings/project_access_tokens.html )")

	// 替换默认的 Usage
	flag.Usage = usage
}

func usage() {
	fmt.Fprintf(os.Stderr, `MR-Tracker Report version: %s
Usage:
	MR-Tracker -v
	MR-Tracker -h
	MR-Tracker -site YOUR_GITLAB_DOMAIN -project NUM -token XXXXXX -start v1.0.0 -end v1.0.1 -post

Options:
`, version)
	flag.PrintDefaults()
}

// TODO
func CheckSettings() error {
	// Site              string
	// Project           int
	// StartTag          string
	// EndTag            string
	// UserRequestVerion bool
	// PostIssue         bool
	// Token             string
	if Settings.Site == "" {
		return fmt.Errorf("site is required")
	}
	if Settings.PostIssue && Settings.Token == "" {
		return fmt.Errorf("token is required for posting issue")
	}
	return nil
}

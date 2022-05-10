package utils

import (
	"flag"
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

// Settings from command line
type UserSettings struct {
	Site      string `yaml:"site"`
	Project   int    `yaml:"project"`
	Branch    string `yaml:"branch"`
	StartTag  string `yaml:"start_tag"`
	EndTag    string `yaml:"end_tag"`
	PostIssue bool   `yaml:"post_issue"`
	Token     string `yaml:"token"`
}

const DefaultConfigFile = ".mr-tracker.yml"

func LoadSettings() (UserSettings, error) {
	if PathExists(DefaultConfigFile) {
		var s UserSettings
		raw := ReadFile(DefaultConfigFile)
		err := yaml.Unmarshal(raw, &s)
		if err != nil {
			return UserSettings{}, err
		}
		return s, nil
	} else {
		fmt.Println("No config file found.")
		return UserSettings{}, nil
	}
}

var (
	RequestVersion bool
	Settings       UserSettings
)

func init() {
	// 注意相同的变量 RequestVersion 绑定了不同的参数，这是允许的
	// 但要确保它们的默认值完全一致，否则会有预期之外的错误
	flag.BoolVar(&RequestVersion, "version", false, "the same as -v")
	flag.BoolVar(&RequestVersion, "v", false, "Print version and exit")

	flag.StringVar(&Settings.StartTag, "start", "", "Set the tag to start analyze(commit excluded)")
	flag.StringVar(&Settings.EndTag, "end", "", "Set the tag to end analyze(commit included)")
	flag.BoolVar(&Settings.PostIssue, "post", false, "Post the result to gitlab issue")

	flag.StringVar(&Settings.Site, "site", "gitlab.com", "Domain of your gitlab instance")
	flag.IntVar(&Settings.Project, "project", 0, "Project ID")
	flag.StringVar(&Settings.Branch, "branch", "", "If you want to track changes other than default branch, set it by this option")
	flag.StringVar(&Settings.Token, "token", "", "Gitlab API token for your project (see: https://docs.gitlab.com/ee/user/project/settings/project_access_tokens.html )")

	// 替换默认的 Usage
	flag.Usage = usage
}

func usage() {
	fmt.Fprintf(os.Stderr, `MR-Tracker Report version: %s
Usage:
	MR-Tracker -v
	MR-Tracker -h
	MR-Tracker -site YOUR_GITLAB_DOMAIN -project YOUR_PROJECT_ID -token GITLAB_API_TOKEN -start v1.0.0 -end v1.0.1 -post

Options:
`, version)
	flag.PrintDefaults()
}

// TODO
func CheckSettings() error {
	if Settings.Site == "" {
		return fmt.Errorf("site is required. run MR-Tracker -h for more information")
	}
	if Settings.Project == 0 {
		return fmt.Errorf("project ID is required. run MR-Tracker -h for more information")
	}
	if Settings.StartTag == "" || Settings.EndTag == "" {
		return fmt.Errorf("start & end tag is required. run MR-Tracker -h for more information")
	}
	if Settings.PostIssue && Settings.Token == "" {
		return fmt.Errorf("token is required for posting issue. run MR-Tracker -h for more information")
	}
	return nil
}

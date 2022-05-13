package utils

import (
	"fmt"
	"runtime"
)

var (
	version   = "v0.1.0"
	commit_id string
	buildTime string
	osArch    string
	go_ver    string
)

// 打印版本信息
func PrintVersion() {
	go_ver = runtime.Version()
	osArch = runtime.GOOS + " " + runtime.GOARCH
	fmt.Println("------------------------------------------")
	fmt.Printf("MR-Tracker Version: %s\nCommit: %s\nBuilt: %s\nOS/Arch: %s\nGo Version: %s\n", version, commit_id, buildTime, osArch, go_ver)
	fmt.Println("------------------------------------------")
	// Log.Infof("MR-Tracker Version: %s\tCommit: %s\tBuilt: %s\tOS/Arch: %s\tGo Version: %s", version, commit_id, buildTime, osArch, go_ver)
}

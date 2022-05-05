package main

import "github.com/MamaShip/MR-Tracker/gitlab"

func main() {
	println("Start Tracking!")
	gitlab.FindMRs("v1.9.3", "v1.9.4")
}
package utils_test

import (
	"net/url"
	"testing"

	"github.com/MamaShip/MR-Tracker/utils"
)

func TestForm(t *testing.T) {
	u := "https://gitlab.example.com/api/v4/projects"
	params := url.Values{}
	params.Set("state", "merged")
	params.Set("order_by", "updated_at")
	params.Set("sort", "desc")
	params.Set("scope", "all")
	params.Set("target_branch", "master")
	result := utils.FormRequest(u, params)
	println(result)
}

type PostIssue struct {
	Title string `json:"title"`
	Body  string `json:"description"`
	Token string `json:"private_token"`
	// Id  int `json:"id"`
}

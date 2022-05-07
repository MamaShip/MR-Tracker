package utils_test

import (
	"encoding/json"
	"fmt"
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
	Body string `json:"description"`
	Token string `json:"private_token"`
	// Id  int `json:"id"`
}

func TestPost(t *testing.T) {
	issue_str :=`## Changes

- [!19](http://gitlab.qitantech.com/playground/projecta/-/merge_requests/19) Update .gitlab-ci.yml by Li Song
- [!4](http://gitlab.qitantech.com/playground/projecta/-/merge_requests/4) 修改模板 by Li Song
- [!3](http://gitlab.qitantech.com/playground/projecta/-/merge_requests/3) 测试 templates by Li Song
- [!2](http://gitlab.qitantech.com/playground/projecta/-/merge_requests/2) Update .gitlab-ci.yml by Li Song`

	i := PostIssue{Title: "Changes of v3.0.1", Body: issue_str, Token: "VhnrgrMbb51t9P3c3ZtG"}
	jsonData, _ := json.Marshal(i)
	resp := utils.Post("http://gitlab.qitantech.com/api/v4/projects/102/issues", jsonData)
	fmt.Println(string(resp))
}

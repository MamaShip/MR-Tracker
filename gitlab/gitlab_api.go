package gitlab

import (
	"fmt"
	"net/url"
)

type Gitlab struct {
	url.URL
	Token string
}

func NewGitlab(project_id int, token string) Gitlab {
	instance := Gitlab{}
	instance.Scheme = "https"
	instance.Host = "gitlab.com"
	instance.Path = fmt.Sprintf("/api/v4/projects/%d", project_id)
	instance.Token = token
	return instance
}

// host=="gitlab.qitantech.com"
func NewCustomGitlab(host string, project_id int, token string) Gitlab {
	instance := Gitlab{}
	instance.Scheme = "http"
	instance.Host = host
	instance.Path = fmt.Sprintf("/api/v4/projects/%d", project_id)
	instance.Token = token
	return instance
}

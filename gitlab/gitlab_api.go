package gitlab

import (
	"net/url"
)

type Gitlab struct {
	url.URL
	Token string
}

func NewGitlab() Gitlab {
	instance := Gitlab{}
	instance.Scheme = "https"
	instance.Host = "gitlab.com"
	instance.Path = "/api/v4/projects/15"
	return instance
}

// host=="gitlab.qitantech.com"
func NewCustomGitlab(host string, token string) Gitlab {
	instance := Gitlab{}
	instance.Scheme = "http"
	instance.Host = host
	instance.Path = "/api/v4/projects/22"
	instance.Token = token
	return instance
}

const GITLAB_API = "https://gitlab.com/api/v4"

func FormUrl() {

}

func FormRequest(url string, params url.Values) string {
	if params == nil {
		return url
	}
	req := url + "?" + params.Encode()
	return req
}

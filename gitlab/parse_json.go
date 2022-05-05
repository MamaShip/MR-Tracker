package gitlab

import "encoding/json"

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	UserName string `json:"username"`
	State    string `json:"state"`
	Url      string `json:"web_url"`
}

type MergeRequest struct {
	Id           int    `json:"id"`
	IId          int    `json:"iid"`
	Title        string `json:"title"`
	State        string `json:"state"`
	MergeUser    User   `json:"merge_user"`
	MergedAt     string `json:"merged_at"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
	SourceBranch string `json:"source_branch"`
	TargetBranch string `json:"target_branch"`
	Author       User   `json:"author"`
	MergeUrl     string `json:"web_url"`
	MergeCommit  string `json:"merge_commit_sha"`
}

func ParseMRs(json_str []byte) []MergeRequest {
	var mrs []MergeRequest
	json.Unmarshal(json_str, &mrs)
	return mrs
}

// GET /projects/:id/repository/branches

type Branch struct {
	Name      string `json:"name"`
	Merged    bool   `json:"merged"`
	Protected bool   `json:"protected"`
	Default   bool   `json:"default"`
	Url       string `json:"web_url"`
}

func ParseBranches(json_str string) []Branch {
	var brs []Branch
	json.Unmarshal([]byte(json_str), &brs)
	return brs
}

// GET /projects/:id/repository/tags

type Tag struct {
	Name      string `json:"name"`
	Message   string `json:"message"`
	Protected bool   `json:"protected"`
	Target    string `json:"target"`
	Url       string `json:"web_url"`
	Commit    Commit `json:"commit"`
}

type Commit struct {
	Id        string   `json:"id"`
	ShortId   string   `json:"short_id"`
	Title     string   `json:"title"`
	CreatedAt string   `json:"created_at"`
	Message   string   `json:"message"`
	ParentIds []string `json:"parent_ids"`
}

func ParseTags(json_str []byte) []Tag {
	var tags []Tag
	json.Unmarshal(json_str, &tags)
	return tags
}

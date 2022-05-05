package gitlab

import (
	"fmt"
	"testing"
)

func TestParse(t *testing.T) {
	json_str := `[{"id":9,"iid":1,"project_id":15,"title":"fix typo","description":"","state":"merged","created_at":"2021-02-25T02:00:30.664Z","updated_at":"2021-02-25T02:00:45.166Z","merged_by":{"id":2,"username":"lisong","name":"Li Song","state":"active","avatar_url":"http://gitlab.qitantech.com/uploads/-/system/user/avatar/2/avatar.png","web_url":"http://gitlab.qitantech.com/lisong"},"merge_user":{"id":2,"username":"lisong","name":"Li Song","state":"active","avatar_url":"http://gitlab.qitantech.com/uploads/-/system/user/avatar/2/avatar.png","web_url":"http://gitlab.qitantech.com/lisong"},"merged_at":"2021-02-25T02:00:45.294Z","closed_by":null,"closed_at":null,"target_branch":"master","source_branch":"debug","user_notes_count":0,"upvotes":0,"downvotes":0,"author":{"id":2,"username":"lisong","name":"Li Song","state":"active","avatar_url":"http://gitlab.qitantech.com/uploads/-/system/user/avatar/2/avatar.png","web_url":"http://gitlab.qitantech.com/lisong"},"assignees":[{"id":2,"username":"lisong","name":"Li Song","state":"active","avatar_url":"http://gitlab.qitantech.com/uploads/-/system/user/avatar/2/avatar.png","web_url":"http://gitlab.qitantech.com/lisong"}],"assignee":{"id":2,"username":"lisong","name":"Li Song","state":"active","avatar_url":"http://gitlab.qitantech.com/uploads/-/system/user/avatar/2/avatar.png","web_url":"http://gitlab.qitantech.com/lisong"},"reviewers":[],"source_project_id":15,"target_project_id":15,"labels":["bug"],"draft":false,"work_in_progress":false,"milestone":null,"merge_when_pipeline_succeeds":false,"merge_status":"can_be_merged","sha":"f958a025ee93c3bb28c63f2f6d73deb92170e66a","merge_commit_sha":"de69c427819b1083e488ebbf5d0b11a2a08d6b4d","squash_commit_sha":null,"discussion_locked":null,"should_remove_source_branch":null,"force_remove_source_branch":true,"reference":"!1","references":{"short":"!1","relative":"!1","full":"lisong/auto-run!1"},"web_url":"http://gitlab.qitantech.com/lisong/auto-run/-/merge_requests/1","time_stats":{"time_estimate":0,"total_time_spent":0,"human_time_estimate":null,"human_total_time_spent":null},"squash":false,"task_completion_status":{"count":0,"completed_count":0},"has_conflicts":false,"blocking_discussions_resolved":true,"approvals_before_merge":null}]`
	mrs := ParseMRs([]byte(json_str))
	fmt.Printf("%+v",mrs)
}

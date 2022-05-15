package gitlab

import (
	"fmt"
	"time"

	"github.com/MamaShip/MR-Tracker/utils"
)

type configBothEnds struct {
	has_start  bool
	start_tag  *Tag
	start_time time.Time
	has_end    bool
	end_tag    *Tag
	end_time   time.Time
}

func configOfBothEnds(start *Tag, end *Tag) (configBothEnds, error) {
	c := configBothEnds{start_tag: start, end_tag: end}
	if start.Name != "" {
		var err error
		c.start_time, err = utils.ParseTime(start.Commit.CreatedAt)
		if err != nil {
			fmt.Println(err)
			return c, err
		}
		c.has_start = true
	}
	if end.Name != "" {
		var err error
		c.end_time, err = utils.ParseTime(end.Commit.CreatedAt)
		if err != nil {
			fmt.Println(err)
			return c, err
		}
		c.has_end = true
	}
	return c, nil
}

// 根据始末 tag 过滤 merge request。
func keepMRsBetween(all_mr []MergeRequest, start Tag, end Tag) []MergeRequest {
	config, err := configOfBothEnds(&start, &end)
	if err != nil {
		return nil
	}

	mrs := make([]MergeRequest, 0, len(all_mr))
	for _, mr := range all_mr {
		// avoid analyze empty MR (too old that no record in gitlab)
		if mr.MergedAt == "" {
			continue
		}
		if belongToRange(&mr, &config) {
			mrs = append(mrs, mr)
		}
	}
	return mrs
}

// 过滤结果剔除了 start tag 指向的 MR。包含了 end tag 指向的 MR。
func belongToRange(mr *MergeRequest, c *configBothEnds) bool {
	if mr.MergeCommit == c.end_tag.Commit.Id {
		return true
	}
	if mr.MergeCommit == c.start_tag.Commit.Id {
		return false
	}

	merge_time, err := utils.ParseTime(mr.MergedAt)
	if err != nil {
		fmt.Printf("Can't sort: %s , ERROR: %s", mr.Title, err)
		return false
	}

	if (c.has_start && merge_time.Before(c.start_time)) ||
		(c.has_end && merge_time.After(c.end_time)) {
		return false
	}
	return true
}

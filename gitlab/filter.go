package gitlab

import (
	"fmt"

	"github.com/MamaShip/MR-Tracker/utils"
)

// 根据始末 tag 过滤 merge request。
// 过滤结果剔除了 start tag 指向的 MR。包含了 end tag 指向的 MR。
func keepMRsBetween(all_mr []MergeRequest, start Tag, end Tag) []MergeRequest {
	start_time, err := utils.ParseTime(start.Commit.CreatedAt)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	end_time, err := utils.ParseTime(end.Commit.CreatedAt)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	mrs := make([]MergeRequest, 0, len(all_mr))
	for _, mr := range all_mr {
		merge_time, err := utils.ParseTime(mr.MergedAt)
		if err != nil {
			fmt.Println(err)
			return nil
		}
		if merge_time.After(start_time) && merge_time.Before(end_time) {
			if mr.MergeCommit != start.Commit.Id {
				mrs = append(mrs, mr)
			}
		} else if mr.MergeCommit == end.Commit.Id {
			mrs = append(mrs, mr)
		}
	}
	return mrs
}

func keepMRsBefore(all_mr []MergeRequest, end Tag) []MergeRequest {
	end_time, err := utils.ParseTime(end.Commit.CreatedAt)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	mrs := make([]MergeRequest, 0, len(all_mr))
	for _, mr := range all_mr {
		merge_time, err := utils.ParseTime(mr.MergedAt)
		if err != nil {
			fmt.Println(err)
			return nil
		}
		if merge_time.Before(end_time) {
			mrs = append(mrs, mr)
		} else if mr.MergeCommit == end.Commit.Id {
			mrs = append(mrs, mr)
		}
	}
	return mrs
}

func keepMRsAfter(all_mr []MergeRequest, start Tag) []MergeRequest {
	start_time, err := utils.ParseTime(start.Commit.CreatedAt)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	mrs := make([]MergeRequest, 0, len(all_mr))
	for _, mr := range all_mr {
		merge_time, err := utils.ParseTime(mr.MergedAt)
		if err != nil {
			fmt.Println(err)
			return nil
		}
		if merge_time.After(start_time) {
			if mr.MergeCommit != start.Commit.Id {
				mrs = append(mrs, mr)
			}
		}
	}
	return mrs
}

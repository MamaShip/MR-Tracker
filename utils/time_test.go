package utils_test

import (
	"testing"
	"time"

	"github.com/MamaShip/MR-Tracker/utils"
)

func TestParseTime(t *testing.T) {
	time_str := "2021-02-25T02:00:45.294Z"
	tt, err := utils.ParseTime(time_str)
	if err != nil {
		t.Errorf("ParseTime(%s) failed with error: %s", time_str, err)
	}
	// fmt.Println(tt)
	if tt.Format(time.RFC3339Nano) != time_str {
		t.Errorf("ParseTime(%s) failed with time: %s", time_str, tt.Format(time.RFC3339))
	}
}

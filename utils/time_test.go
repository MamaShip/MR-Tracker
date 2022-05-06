package utils_test

import (
	"testing"
	"time"

	"github.com/MamaShip/MR-Tracker/utils"
	"github.com/stretchr/testify/assert"
)

func TestParseTime(t *testing.T) {
	time_str := "2021-02-25T02:00:45.294Z"
	tt, err := utils.ParseTime(time_str)
	if err != nil {
		t.Errorf("ParseTime(%s) failed with error: %s", time_str, err)
	}
	// fmt.Println(tt)
	assert.Equal(t, tt.Format(time.RFC3339Nano), time_str)
}

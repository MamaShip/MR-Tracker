package gitlab

import (
	"fmt"
	"testing"

	"github.com/MamaShip/MR-Tracker/utils"
	"github.com/stretchr/testify/assert"
)

func TestGetLatestTag(t *testing.T) {
	var s utils.UserSettings
	s.Site = "gitlab.com"
	s.Project = 31285645
	s.Latest = "v1.0.2"

	start, end := GetLatestTag(s)
	fmt.Println(start, end)
	assert.Equal(t, s.Latest, end)
	assert.Equal(t, "v1.0.1", start)
}

func Test_getSemver(t *testing.T) {
	v, err := getSemver("v2.3.4_pre")
	assert.Nil(t, err)
	fmt.Println(v)
	assert.Equal(t, uint64(2), v.Major)
	assert.Equal(t, uint64(3), v.Minor)
	assert.Equal(t, uint64(4), v.Patch)
}

func Test_isFormalVersion(t *testing.T) {
	assert.True(t, isFormalVersion("v1.0.0"))
	assert.True(t, isFormalVersion("V1.0.0"))
	assert.True(t, isFormalVersion("3.0.110"))
	assert.False(t, isFormalVersion("v1.0.0_pre"))
	assert.False(t, isFormalVersion("v1.0.0x"))
	assert.False(t, isFormalVersion("vv1.0.0"))
}

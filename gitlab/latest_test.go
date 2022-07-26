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

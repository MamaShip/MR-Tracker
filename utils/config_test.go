package utils_test

import (
	"fmt"
	"testing"

	"github.com/MamaShip/MR-Tracker/utils"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
)

func TestSaveSettings(t *testing.T) {
	var s utils.UserSettings
	s.Site = "gitlab.com"
	s.Branch = "master"
	s.Project = 102
	s.Token = "xxxxxxxxxxxxxxxxx"
	s.PostIssue = true
	s.StartTag = "v3.0.1"
	s.EndTag = "v3.0.2"
	bytes, err := yaml.Marshal(s)
	assert.Nil(t, err)
	fmt.Println(string(bytes))
}

func TestLoadSettings(t *testing.T) {
	s, err := utils.LoadSettings()
	assert.Nil(t, err)
	fmt.Printf("%+v\n", s)
}

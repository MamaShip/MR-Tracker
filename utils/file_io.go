package utils

import (
	"io/ioutil"
	"os"
)

// read whole file content
func ReadFile(filepath string) []byte {
	// Log.Debug("Read All: ", filepath)
	content, err := ioutil.ReadFile(filepath)
	if err != nil {
		panic(err)
	}
	return content
}

func PathExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	} else if err != nil {
		panic(err)
	}
	return true
}

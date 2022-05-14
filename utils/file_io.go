package utils

import (
	"fmt"
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

func Write2File(changes string, file string) error {
	output_file, err := os.Create(file)
	if err != nil {
		fmt.Println("failed to open outpout file:", output_file, err)
		return err
	}
	defer output_file.Close()
	_, err = output_file.WriteString(changes)
	return err
}

package utils

import (
	"io/ioutil"
	"log"
	"net/http"
)

func Get(url string) []byte {
	// resp, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	if resp.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", resp.StatusCode, body)
	}
	//Convert the body to type string
	// sb := string(body)
	return body
}

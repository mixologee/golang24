package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"net/http"
	"log"
)

func main() {
	// create the data object for the POST
	postData := strings.NewReader(`{ "some": "json"}`)
	// run the POST and store in response
	response, err := http.Post("https://httpbin.org/post", "application/json", postData)
	
	// handle errors
	if err != nil {
		log.Fatal(err)
	}

	// close the request when done
	defer response.Body.Close()

	// read the request
	body, err := ioutil.ReadAll(response.Body)

	// handle errors
	if err != nil {
		log.Fatal(err)
	}

	// print the info
	fmt.Printf("%s", body)
}

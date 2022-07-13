package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// create your response variable to hold your GET request data
	response, err := http.Get("https://ipchicken.com/")

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

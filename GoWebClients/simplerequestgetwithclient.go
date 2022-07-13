package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// create the client object
	client := &http.Client{}
	// create your response variable to hold your GET request data
	request, err := http.NewRequest("GET", "https://ipchicken.com/", nil)

	// handle errors
	if err != nil {
		log.Fatal(err)
	}
	// process the reuqest using the client
	response, err := client.Do(request)
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

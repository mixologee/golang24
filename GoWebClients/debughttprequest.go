package main

import (
	"fmt"
	"os"
	"log"
	"net/http"
	"net/http/httputil"
	"io/ioutil"
)

func main() {
	// setup variables to be used
	debug := os.Getenv("DEBUG")
	client := &http.Client{}
	request, err := http.NewRequest("GET", "https://ipchicken.com/", nil)

	// handle errors
	if err != nil {
		log.Fatal(err)
	}

	if debug == "1" {
		debugRequest, err := httputil.DumpRequestOut(request, true)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s", debugRequest)
	}
	// process the reuqest using the client
	response, err := client.Do(request)
	// close the request when done
	defer response.Body.Close()

	if debug == "1" {
		debugResponse, err := httputil.DumpResponse(response, true)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s", debugResponse)
	}

	// read the request
	body, err := ioutil.ReadAll(response.Body)

	// handle errors
	if err != nil {
		log.Fatal(err)
	}

	// print the info
	fmt.Printf("%s", body)
}

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// struct to hold the data that is returned
type User struct {
	Name string `json:"name"`
	Blog string `json:"blog"`
}

func main() {
	// set your variables
	var u User
	res, err := http.Get("https://api.github.com/users/shapeshed")
	// always error handle
	if err != nil {
		log.Fatal(err)
	}
	// close the response when we are done
	defer res.Body.Close()
	// decode the json data
	err = json.NewDecoder(res.Body).Decode(&u)
	// always error handle
	if err != nil {
		log.Fatal(err)
	}
	// print out your info
	fmt.Printf("%+v\n", u)
}
package main

import (
	"fmt"
	"encoding/json"
	"log"
)

// struct for your new variable type Person
// tag for json variables so not capitalized when doing json encoding
// omitempty ignores empty fields when encoding to json format
type Person struct {
	Name    string		`json:"name, omitempty"`
	Age     int			`json:"age, omitempty"`
	Hobbies []string	`json:"hobbies, omitempty"`
}

func main() {
	// create variable full of hobbies
	hobbies := []string{"Cycling", "Cheese", "Techno"}
	// create a person
	p := Person{
		Name:    "George",
		Age:     40,
		Hobbies: hobbies,
	}
	// pring values
	fmt.Printf("%+v\n", p)

	// create variable for encoded data
	jsonByteData, err := json.Marshal(p)

	// handle errors
	if err != nil {
		log.Fatal(err)
	}
	// convert the data to a string and print
	jsonStringData := string(jsonByteData)
	fmt.Println(jsonStringData)
}
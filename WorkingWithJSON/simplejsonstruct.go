package main

import "fmt"

// struct for your new variable type Person
type Person struct {
	Name    string
	Age     int
	Hobbies []string
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
}
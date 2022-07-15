package main

import (
	"fmt"
	"log"
  // external library to handle TOML format
	"github.com/BurntSushi/toml"
)

// struct to hold config data
type Config struct {
	Name   string
	Awake  bool
	Hungry bool
}

func main() {
  // create the variable to hold data
	c := Config{}
  // decode the file data and put in variable
	_, err := toml.DecodeFile("config.toml", &c)
  // error handle
	if err != nil {
		log.Fatal(err)
	}
  // prit out data
	fmt.Printf("%+v\n", c)
}

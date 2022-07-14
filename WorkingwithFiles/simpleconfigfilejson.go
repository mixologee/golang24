package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

// struct to hodl the config data with json tags
type Config struct {
	Name   string `json:"name"`
	Awake  bool   `json:"awake"`
	Hungry bool   `json:"hungry"`
}

func main() {
  // variable to hold the read data
	f, err := ioutil.ReadFile("config.json")
  // error handle
	if err != nil {
		log.Fatal(err)
	}
  // create the variable to hold the file data
	c := Config{}
  // convert the data from json format and dump into Config variable
	err = json.Unmarshal(f, &c)
  //error handle
	if err != nil {
		log.Fatal(err)
	}
  // print values
	fmt.Printf("%+v\n", c)
}

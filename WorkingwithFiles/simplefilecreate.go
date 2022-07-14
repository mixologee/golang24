package main

import (
	"io/ioutil"
	"log"
)

func main() {
	// create your byte variable to hold data to be put in the file
	b := make([]byte, 0)
	// write to the file (file name, content, permissions)
	err := ioutil.WriteFile("myfile.txt", b, 0644)
	// error handle
	if err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"log"
	"os"
)

func main() {
  // create the variable to store the delete results
	err := os.Remove("./deleteme.txt")
  // handle error
	if err != nil {
		log.Fatal(err)
	}
}

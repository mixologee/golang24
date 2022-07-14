package main

import (
        "io/ioutil"
        "log"
        "fmt"
)

func main() {
        // create your reader variable
        files, err := ioutil.ReadDir(".")
        // error handle
        if err != nil {
                log.Fatal(err)
        }
        // loop through and list info
        for _, file := range files {
                fmt.Println(file.Mode(), file.Name())
        }

}

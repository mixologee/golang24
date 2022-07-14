package main

import (
        "io/ioutil"
        "log"
        "fmt"
)

func main() {
        // create your variable to hold data to be put in the file
        s := "Hello File!"
        // write to the file (file name, content, permissions)
        err := ioutil.WriteFile("myfile.txt", i[]byte(s), 0644)
        // error handle
        if err != nil {
                log.Fatal(err)
        }
}

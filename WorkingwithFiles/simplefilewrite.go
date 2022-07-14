package main

import (
        "io/ioutil"
        "log"
)

func main() {
        // create your variable to hold data to be put in the file
        s := "Hello File!\n"
        // write to the file (file name, content, permissions)
        err := ioutil.WriteFile("myfile.txt", []byte(s), 0644)
        // error handle
        if err != nil {
                log.Fatal(err)
        }
}

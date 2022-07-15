package main

import (
        "fmt"
        "io/ioutil"
        "log"
)

func main() {
        // create file reader variable
        fileBytes, err := ioutil.ReadFile("/proc/loadavg")
        if err != nil {
                log.Fatal(err)
        }
        // print raw data
        fmt.Println(fileBytes)
        // convert to string and print data
        fileString := string(fileBytes)
        fmt.Println(fileString)
}

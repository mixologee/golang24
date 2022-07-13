package main

import (
        "fmt"
        "os"
)

func main() {
        // loop through and print out arguments that are passed from commandline
        // Test: go run simplerawarguments.go arg1 arg2 arg3
        for i, arg := range os.Args {
                fmt.Println("argument", i, "equals", arg)
        }
}

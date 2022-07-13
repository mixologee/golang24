package main

import (
        "flag"
        "fmt"
)

func main() {
        // create a variable that is tied to a flag
        // format is flag name, default value, help string
        s := flag.String("s", "Hello World", "Some kind of help text")
        // parses the flags
        flag.Parse()
        // prints the value of s
        fmt.Println("s equals: ", *s)
}

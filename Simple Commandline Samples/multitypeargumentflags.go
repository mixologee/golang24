package main

import (
        "flag"
        "fmt"
)

func main() {
        // create different variables of different types and assign to a flag
        // format is flag name, default value, help string
        s := flag.String("s", "Hello World", "This needs to be a string. (Default: Hello World)")
        i := flag.Int("i", 1, "This needs to be an integer. (Default: 1)")
        b := flag.Bool("b", true, "This needs to be a boolean. (Default: false)")
        // parses the flags
        flag.Parse()
        // prints the values of your variables
        fmt.Println("s equals:", *s)
        fmt.Println("i equals:", *i)
        fmt.Println("b equals:", *b)
}

package main

import (
        "fmt"
)

type Animal struct {
        Name    string
        Color   string
}

func main() {
        // printing a string
        s := "Hello World"
        fmt.Printf("string is %v\n", s)

        // printing a struct
        a := Animal {
                Name: "Cat",
                Color: "black",
        }
        // check the outputs make sure you get what you need
        fmt.Printf("%v\n", a)
        fmt.Printf("%+v\n", a)
}

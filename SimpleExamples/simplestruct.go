package main

import (
	"fmt"
)

type Drink struct {
	Name	string
	Ice	bool
}

func main() {
	a := Drink{
		Name: "Lemonade",
		Ice: true,
	}
	b := a
	b.Ice = false
	fmt.Printf("%+v",a)
	fmt.Printf("%+v",b)
	fmt.Printf("%p",&a)
	fmt.Printf("%p",&b)

	c := &a
	c.Ice = false
	fmt.Printf("%+v",a)
        fmt.Printf("%+v",*c)
        fmt.Printf("%p",&a)
        fmt.Printf("%p",c)
}

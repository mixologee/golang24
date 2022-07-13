package main

import (
	"fmt"
)

func callDefer(x string){
	fmt.Println("deferred " + x)
}

func main () {

	defer fmt.Println("1st defer")
	defer fmt.Println("2nd defer")
	for i:=0;i<5;i++ {
		defer fmt.Println(i)
	}
	callDefer("1")
	defer fmt.Println("3rd defer")
	defer callDefer("2")
}

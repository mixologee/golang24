package main

import (
	"fmt"
	"time"
)

func slowFunc() {
	time.Sleep(time.Second * 2)
	fmt.Println("sleeper() finished")
}

func main() {
	go slowFunc()
	fmt.Println("I am now printed before slowFunc() is done.")
	// lets wait for slowFunc() to complete
	time.Sleep(time.Second * 3)
}

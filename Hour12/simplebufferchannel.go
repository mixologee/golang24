package main

import (
	"fmt"
	"time"
)

func receiver(c chan string) {
	for msg := range c {
		fmt.Println(msg)
	}
}

func main () {
	messages := make(chan string, 2)
	messages <- "hello"
	messages <- "world"
	close(messages)
	fmt.Println("Pushed two messages then close the channel with no receiver.")
	time.Sleep(time.Second * 1)
	receiver(messages)
}
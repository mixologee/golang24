package main

import (
	"fmt"
	"time"
)

func pinger(c chan string) {
	t := time.NewTicker(1 * time.Second)

	for { 
		c <- "ping"
		<-t.C // drain the timer if it is not stopped
	}
}

func main () {
	messages := make(chan string)
	go pinger(messages)
	for {
		msg := <-messages
		fmt.Println(msg)
	}
}
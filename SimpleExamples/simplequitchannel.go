package main

import (
	"fmt"
	"time"
)

func sender(c chan string) {
	t := time.NewTicker(1 * time.Second)

	for {
		c <- "Message being sent."
		<-t.C // drain the timer if it is not stopped
	}
}

func main () {
	messages := make(chan string)
	stop := make(chan bool)
	go sender(messages)
	go func() {
		time.Sleep(time.Second*2)
		fmt.Println("Time's up!")
		stop <- true
	}()

	for {
		select {
		case <-stop:
			return
		case msg := <-messages:
			fmt.Println(msg)
		}
	}
}
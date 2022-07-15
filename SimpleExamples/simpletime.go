package main

import (
	"fmt"
	"time"
	"log"
)

func main() {
	// print the current time
	fmt.Println(time.Now())

	// sleep example
	fmt.Println("I'm sleeping")
	time.Sleep(3 * time.Second)
	fmt.Println("I'm awake")

	// example of using a ticker 
	c := time.Tick(5 * time.Second)
	i := 0
	for t := range c {
		fmt.Printf("The time is now %v\n", t)
		//to make this tick continuously, remove the counter code involving i
		i += 1
		if i > 5 {
			break
		}
	}

	// example of parsing time from strings
	s := "1978-01-31T15:04:05+07:00"
	t, err := time.Parse(time.RFC3339, s)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(t)
	// example of Time struct properties
	fmt.Printf("The hour is %v\n", t.Hour())
	fmt.Printf("The minute is %v\n", t.Minute())
	fmt.Printf("The second is %v\n", t.Second())
	fmt.Printf("The day is %v\n", t.Day())
	fmt.Printf("The month is %v\n", t.Month())
	fmt.Printf("UNIX time is %v\n", t.Unix())
	fmt.Printf("The day of the week is %v\n", t.Weekday())

	// example of Time math
	s1 := "2017-09-03T18:00:00+00:00"
	s2 := "2017-09-04T18:00:00+00:00"
	today, err := time.Parse(time.RFC3339, s1)
	if err != nil {
		log.Fatal(err)
	}
	tomorrow, err := time.Parse(time.RFC3339, s2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(today.After(tomorrow))
	fmt.Println(today.Before(tomorrow))
	fmt.Println(today.Equal(tomorrow))

	// example of setting up a timeout with a for loop
	fmt.Println("You have two seconds to calculate 19 * 4")
	for {
		select {
		// create the switch case to wait for 2 seconds on the timer
		case <-time.After(2 * time.Second):
			fmt.Println("Time's up! The answer is 74. Did you get it?")			
			return
		}
	}
}
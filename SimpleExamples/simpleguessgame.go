package main

import (
        "fmt"
        "bufio"
        "strings"
        "os"
)

func main () {
        // create input reader
        reader := bufio.NewReader(os.Stdin)
        fmt.Println("Guess my pet's name: ")
        // read what is typed until enter is pressed
        text, _ := reader.ReadString('\n')
        // remove the carriage return from the text
        text = strings.Replace(text,"\n","",-1)
        // print out their text for debugging
        fmt.Println("[DEBUG] User entered: ",text)
        if text == "Buster" {
                fmt.Println("Correct!")
        } else {
                fmt.Println("Sorry, you lose.")
        }
}

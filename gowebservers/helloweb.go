package main

import (
        "net/http"
)

// function to handle web requests
func helloWeb(w http.ResponseWriter, r *http.Request) {
        // response to web request
        w.Write([]byte("Hello Web!\r\n"))
}

// main function
func main() {
        // assign the handler function to a path
        http.HandleFunc("/", helloWeb)
        // start listening for web requests on a specific port
        http.ListenAndServe(":8000", nil)
}

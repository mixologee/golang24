package main

import (
        "net/http"
)

func helloWeb(w http.ResponseWriter, r *http.Request) {
        // check to see if the path is anything but /, if so throw 404 error
        if r.URL.Path != "/" {
                http.NotFound(w,r)
                return
        }
        // response to web request
        w.Write([]byte("Hello Web!\r\n"))
}

func main() {
        // assign the handler function to a path
        http.HandleFunc("/", helloWeb)
        // start listening for web requests on a specific port
        http.ListenAndServe(":8000", nil)
}

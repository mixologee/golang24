package main

import (
        "net/http"
)

func helloWeb(w http.ResponseWriter, r *http.Request) {
        // send 404 for anything beyond the landing page
        if r.URL.Path != "/" {
                http.NotFound(w,r)
                return
        }

        switch r.Header.Get("Accept") {
                case "application/json":
                        // set a content type in the header
                        w.Header().Set("Content-Type", "application/json; characterset=utf-8")
                        // response to web request
                        w.Write([]byte(`{"message":"Hello Web!\r\n"}`))
                case "application/xml":
                        // set a content type in the header
                        w.Header().Set("Content-Type", "application/xml; characterset=utf-8")
                        // response to web request
                        w.Write([]byte(`<?xml version="1.0" encoding="utf-8"?><Message>Hello Web!</Message>`))
                default:
                        // set a content type in the header
                        w.Header().Set("Content-Type", "text/plain; characterset=utf-8")
                        // response to web request
                        w.Write([]byte("Hello Web!\r\n"))
        }
}

func main() {
        // assign the handler function to a path
        http.HandleFunc("/", helloWeb)
        // start listening for web requests on a specific port
        http.ListenAndServe(":8000", nil)
}

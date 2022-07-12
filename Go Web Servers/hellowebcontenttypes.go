package main

import (
        "net/http"
)

func helloWeb(w http.ResponseWriter, r *http.Request) {
        if r.URL.Path != "/" {
                http.NotFound(w,r)
                return
        }

        switch r.Header.Get("Accept") {
                case "application/json":
                         w.Header().Set("Content-Type", "application/json; characterset=utf-8")
                         w.Write([]byte(`{"message":"Hello Web!\r\n"}`))
                case "application/xml":
                         w.Header().Set("Content-Type", "application/xml; characterset=utf-8")
                         w.Write([]byte(`<?xml version="1.0" encoding="utf-8"?><Message>Hello Web!</Message>`))
                default:
                        w.Header().Set("Content-Type", "text/plain; characterset=utf-8")
                        w.Write([]byte("Hello Web!\r\n"))
        }
}

func main() {
        http.HandleFunc("/", helloWeb)
        http.ListenAndServe(":8000", nil)
}

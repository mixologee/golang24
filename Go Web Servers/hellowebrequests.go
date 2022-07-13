package main

import (
        "net/http"
)

func helloWeb(w http.ResponseWriter, r *http.Request) {
        if r.URL.Path != "/" {
                http.NotFound(w,r)
                return
        }

        switch r.Method {
                case "GET":
                         w.Write([]byte("This was a GET request.\n"))
                case "POST":
                         w.Write([]byte("This was a POST request.\n"))
                default:
                        w.WriteHeader(http.StatusNotImplemented)
                        w.Write([]byte(http.StatusText(http.StatusNotImplemented) + "\n"))
        }
}

func main() {
        http.HandleFunc("/", helloWeb)
        http.ListenAndServe(":8000", nil)
}

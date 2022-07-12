package main

import (
        "net/http"
)

func helloWeb(w http.ResponseWriter, r *http.Request) {
        if r.URL.Path != "/" {
                http.NotFound(w,r)
                return
        }
        w.Write([]byte("Hello Web!\r\n"))
}

func main() {
        http.HandleFunc("/", helloWeb)
        http.ListenAndServe(":8000", nil)
}

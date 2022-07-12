package main

import (
        "net/http"
        "fmt"
        "io/ioutil"
        "log"
)

func helloWeb(w http.ResponseWriter, r *http.Request) {
        if r.URL.Path != "/" {
                http.NotFound(w,r)
                return
        }

        switch r.Method {
                case "GET":
                        for k, v := range r.URL.Query() {
                                fmt.Printf("s: %s\n", k, v)
                        }
                        w.Write([]byte("This was a GET request.\n"))
                case "POST":
                        reqBody, err := ioutil.ReadAll(r.Body)
                        if err != nil {
                                log.Fatal(err)
                        }

                        fmt.Printf("%s\n", reqBody)
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

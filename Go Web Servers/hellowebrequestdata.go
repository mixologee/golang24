package main

import (
        "net/http"
        "fmt"
        "io/ioutil"
        "log"
)

func helloWeb(w http.ResponseWriter, r *http.Request) {
        // send 404 for anything beyond the landing page
        if r.URL.Path != "/" {
                http.NotFound(w,r)
                return
        }
        
        // while not handled here currently, any data beign sent or received to web servers should be scrubbed prior to use
        // assume all data being sent to your web page is malicious and code accordingly
        switch r.Method {
                case "GET":
                        // read in the request info
                        for k, v := range r.URL.Query() {
                                fmt.Printf("s: %s\n", k, v)
                        }
                        // response to web request
                        w.Write([]byte("This was a GET request.\n"))
                case "POST":
                        // read in the data being passed
                        reqBody, err := ioutil.ReadAll(r.Body)
                        // print any error info if required
                        if err != nil {
                                log.Fatal(err)
                        }
                        // print out the data locally
                        fmt.Printf("%s\n", reqBody)
                        // response to web request
                        w.Write([]byte("This was a POST request.\n"))
                default:
                        // set header value
                        w.WriteHeader(http.StatusNotImplemented)
                        // response to web request
                        w.Write([]byte(http.StatusText(http.StatusNotImplemented) + "\n"))
        }
}

func main() {
        // assign the handler function to a path
        http.HandleFunc("/", helloWeb)
        // start listening for web requests on a specific port
        http.ListenAndServe(":8000", nil)
}

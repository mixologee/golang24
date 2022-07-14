package main

import (
        "io"
        "log"
        "os"
)

func main() {
        // create variable of source file
        from, err := os.Open("./myfile.txt")
        if err != nil {
                log.Fatal(err)
        }
        // close source file when done
        defer from.Close()

        // create the variable for the copied file
        to, err := os.OpenFile("./myfile.copy.txt", os.O_RDWR|os.O_CREATE, 0666)
        // erro handle
        if err != nil {
                log.Fatal(err)
        }
        // close the file when we are done
        defer to.Close()
        // copy file
        _, err = io.Copy(to, from)
        // handle errors
        if err != nil {
                log.Fatal(err)
        }
}

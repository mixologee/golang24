package main

import (
        "log"
        "errors"
        "os"
)

func LogToFile(fname string) {
        // open the file, create if needed, append if it exists
        f, err := os.OpenFile(fname, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
        // handle errors
        if err != nil {
                log.Fatal(err)
        }
        // make sure we close the file when we are done
        defer f.Close()
        // set the target for log messages
        log.SetOutput(f)
        // send messages to the log target
        for i := 1;i <=5;i++ {
                log.Printf("log iterate %d", i)
        }
}

func main() {

        // log to screen
        log.Printf("Simple log to output message")

        // logging to a file
        LogToFile("test-log.log")

        // logging fatal error
        var errFatal = errors.New("oof, we just crashed")
        log.Fatal(errFatal)
}

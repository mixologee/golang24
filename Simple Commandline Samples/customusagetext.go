package main

import (
        "flag"
        "fmt"
        "os"
)

func main() {
        // sample of how to create a custom usage message
        flag.Usage = func(){
                usageText := `Usage customusage text [OPTION]
                Example of custom usage output text

                -s, --s         This needs to be a string. (Default: Hello World)
                -i, --i         This needs to be an integer. (Default: 1)
                -b, --b         Using this flag changes boolean value. (Default: false)`
                fmt.Fprintf(os.Stderr, "%s\n", usageText)
        }

        // create a variable that is tied to a flag
        // format is flag name, default value, help string
        s := flag.String("s", "Hello World", "This needs to be a string. (Default: Hello World)")
        i := flag.Int("i", 1, "This needs to be an integer. (Default: 1)")
        b := flag.Bool("b", false, "Using this flag changes boolean value. (Default: false)")
        // parses the flags
        flag.Parse()
        // prints the value of s
        fmt.Println("s equals:", *s)
        fmt.Println("i equals:", *i)
        fmt.Println("b equals:", *b)
}

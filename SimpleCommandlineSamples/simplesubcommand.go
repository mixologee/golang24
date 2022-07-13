package main

import (
        "flag"
        "fmt"
        "os"
        "strings"
)

func flagUsage() {

 usageText := `This is a simple subcommand example
                Usage: simplesubcommand command [arguments]

                Commands are:
                uppercase       sets a string to all uppercase
                lowercase       sets a string to all lowercase

                Use "simplesubcommand [command] --help" for more information `

                fmt.Fprintf(os.Stderr, "%s\n", usageText)
}

func main() {
        // set the usage text to our flag usage function
        flag.Usage = flagUsage
        // create a variable that is tied to a flag
        upperCmd := flag.NewFlagSet("uppercase", flag.ExitOnError)
        lowerCmd := flag.NewFlagSet("lowercase", flag.ExitOnError)
        // check if they need general help or command help
        if len(os.Args) == 1 {
                flag.Usage()
                return
        }
        // format is flag name, default value, help string
        switch os.Args[1] {
                case "uppercase":
                        s := upperCmd.String("s", "", "This needs to be a string.")
                        upperCmd.Parse(os.Args[2:])
                        fmt.Println(strings.ToUpper(*s))
                case "lowercase":
                        s := lowerCmd.String("s", "", "This needs to be a string.")
                        lowerCmd.Parse(os.Args[2:])
                        fmt.Println(strings.ToLower(*s))
                default:
                        flag.Usage()
        }
}

package simpletest3

func translate(s string) string {
        switch s {
                case "en-US":
                        return "Hello " + s
                case "fr-FR":
                        return "Bonjour " + s
                case "it-IT":
                        return "Ciao " + s
                default:
                        return "Hello " + s
        }
}

func Greeting(s string) string {
        return "Hello " + s
}

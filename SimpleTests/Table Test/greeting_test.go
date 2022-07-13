package simpletest3

import "testing"

type GreetingTest struct {
        name    string
        locale  string
        want    string
}

var greetingTests = []GreetingTest {
        {"Name", "en-US", "Hello Name"},
        {"Name", "fr-FR", "Bonjour Name"},
        {"Name", "it-IT", "Ciao Name"},
}

func TestGreeting(t *testing.T) {
        for _, test := range greetingTests {
                got := Greeting(test.name, test.locale)
                if got != test.want {
                        t.Errorf("Greeting(%s,%s) gave %v instead of %v", test.name, test.locale, got, test.want)
                }
        }
}

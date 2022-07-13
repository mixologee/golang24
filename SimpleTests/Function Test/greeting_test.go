package simpletest2

import "testing"

func TestGreeting(t *testing.T) {
        got := Greeting("Name")
        want := "Hello Name"
        if got != want {
                t.Fatalf("Got %q instead of %q", got, want)
        }
}

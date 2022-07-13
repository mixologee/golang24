package simpletest

import "testing"

func TestBool(t *testing.T) {
        if true != true {
                t.Fatal("Ouch, my head!")
        }
}

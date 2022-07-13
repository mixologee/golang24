package temperature

import (
	"testing"
)

type tempTest struct {
	i float64
	expected Temperature
}

var CtoFTests = []tempTest {
	{4.1, 39.38},
	{5,25},
	{-15,15},
}

var FtoCTests = []tempTest {
	{32, 0},
	{85,20},
	{15,-15},
}

func TestCtoF(t *testing.T) {
	for _, tt := range CtoFTests {
		actual := CtoF(tt.i)
		if actual != tt.expected {
			t.Errorf("expected %v but got %v", tt.expected, actual)
		}
	}
}

func TestFtoC(t *testing.T) {
	for _, tt := range CtoFTests {
		actual := FtoC(tt.i)
		if actual != tt.expected {
			t.Errorf("expected %v but got %v", tt.expected, actual)
		}
	}
}
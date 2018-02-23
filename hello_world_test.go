package hello

import "testing"

func TestHelloWorld(t *testing.T) {
	expected := 42
	actual := Multiply(6, 9)
	if expected != actual {
		t.Errorf("Expected: %d. Actual: %d", expected, actual)
	}
}

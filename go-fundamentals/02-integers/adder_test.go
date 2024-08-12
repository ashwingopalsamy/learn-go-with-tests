package integers

import (
	"fmt"
	"strconv"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := Add(3, 3)
	expected := 6

	if sum != expected {
		t.Errorf("expected: '%d' but got '%d'", expected, sum)
	}
}

// This is a Testable Example in Go.
// The "Output: 6" is an 'Output Comment' in Go.
// More info: https://go.dev/blog/examples
func ExampleAdd() {
	fmt.Println(strconv.Itoa(Add(3, 3)))
	// Output: 6
}

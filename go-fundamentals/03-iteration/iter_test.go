package iteration

import (
	"fmt"
	"testing"
)

func TestIter(t *testing.T) {
	repeated := iter("a")
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("Expected: '%q', Got: '%q'", expected, repeated)
	}
}

func ExampleIter() {
	repeated := iter("a")
	fmt.Println(repeated)
	// Output: aaaaa
}

func BenchmarkIter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		iter("a")
	}
}

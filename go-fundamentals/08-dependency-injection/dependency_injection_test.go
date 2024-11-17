package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	t.Run("test", func(t *testing.T) {
		buffer := bytes.Buffer{}
		Greet(&buffer, "Ashwin")

		got := buffer.String()
		want := "Hello, Ashwin"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}

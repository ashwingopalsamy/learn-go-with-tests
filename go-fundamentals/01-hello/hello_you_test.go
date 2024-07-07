package main

import "testing"

func TestHelloYou(t *testing.T) {
	got := HelloName("Ashwin")
	want := "Hello, Ashwin! 👋"
	if got != want {
		t.Errorf("want: %q, got: %q", want, got)
	}
}

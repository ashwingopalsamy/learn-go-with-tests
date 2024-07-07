package main

import "testing"

func TestHelloYou(t *testing.T) {
	t.Run("Saying Hello to People", func(t *testing.T) {
		got := HelloName("Ashwin")
		want := "Hello, Ashwin! 👋"
		if got != want {
			t.Errorf("want: %q, got: %q", want, got)
		}
	})

	t.Run("Saying Hello to World", func(t *testing.T) {
		got := HelloName("")
		want := "Hello, World! 👋"
		if got != want {
			t.Errorf("want: %q, got:%q", want, got)
		}
	})

}

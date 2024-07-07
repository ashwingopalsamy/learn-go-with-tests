package main

import "testing"

func TestHelloYou(t *testing.T) {
	t.Run("Saying Hello to People", func(t *testing.T) {
		got := HelloName("Ashwin", "English")
		want := "Hello, Ashwin! 👋"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Saying Hello to World", func(t *testing.T) {
		got := HelloName("", "")
		want := "Hello, World! 👋"
		assertCorrectMessage(t, got, want)
	})

	t.Run("In Spanish, say hello to people", func(t *testing.T) {
		got := HelloName("Elodie", "Spanish")
		want := "Hola, Elodie! 👋"
		assertCorrectMessage(t, got, want)
	})

	t.Run("In French, say hello to people", func(t *testing.T) {
		got := HelloName("Mira", "French")
		want := "Bonjour, Mira! 👋"
		assertCorrectMessage(t, got, want)
	})

}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("want: %q, got:%q", want, got)
	}
}

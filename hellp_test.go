package main

import "testing"

func TestHellp(t *testing.T) {
	t.Run("syaing hellp to people", func(t *testing.T) {
		got := Hellp("Josh")
		want := "Hellp, Josh!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("syaing 'Hellp, World!' when an empty string is supplied", func(t *testing.T) {
		got := Hellp("")
		want := "Hellp, World!"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

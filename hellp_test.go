package main

import "testing"

func TestHellp(t *testing.T) {
	got := Hellp("Josh")
	want := "Hellp, Josh!"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

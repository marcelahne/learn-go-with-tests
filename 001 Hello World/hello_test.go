package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Marcel")
	want := "Hello, Marcel"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

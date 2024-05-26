package helloworld

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Juan")
	want := "Hello, Juan"

	if got != want {
		t.Fatalf("got %q, want %q", got, want)
	}
}

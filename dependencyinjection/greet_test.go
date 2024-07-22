package dependencyinjection

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	var buffer bytes.Buffer
	Greet(&buffer, "John")

	got := buffer.String()
	want := "Hello, John"

	if got != want {
		t.Fatalf("got %q, want %q", got, want)
	}
}

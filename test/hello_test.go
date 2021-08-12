package hello

import (
	"testing"
	"my.app/pkg"
)

func TestHello(t *testing.T) {
	want := "Hello, world"
	if got := hello.Hello(); got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}
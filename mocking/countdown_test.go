package main

import (
	"bytes"
	"fmt"
	"testing"
)

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}

	Countdown(buffer)

	got := buffer.String()
	want := `3
2
1
Go!`

	fmt.Println("the countdown string", want)
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

package main

import (
	"fmt"
	"io"
	"time"
)

const finalWord = "Go!"
const countdownStart = 3

type Sleeper interface {
	Sleep()
}

type DefaultSleeper struct{}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		sleeper.Sleep()
		// time.Sleep(1 * time.Second)
		fmt.Fprintln(out, i)
	}

	sleeper.Sleep()
	// time.Sleep(1 * time.Second)
	fmt.Fprint(out, finalWord)
}

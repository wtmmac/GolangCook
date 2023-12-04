package main

import (
	"fmt"
	"io"
	"os"
)

const (
	finalWord      = "Go!"
	countdownStart = 3
)

func Countdown(out io.Writer) {
	for l := countdownStart; l > 0; l-- {
		fmt.Fprintln(out, l)
	}
	fmt.Fprint(out, finalWord)
}

func main() {
	Countdown(os.Stdout)
}

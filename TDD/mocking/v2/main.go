package main

import (
	"fmt"
	"io"
	"os"
)

const (
	finalWorld     = "Go!"
	countdownStart = 3
)

func Countdown(out io.Writer) {
	for l := countdownStart; l > 0; l-- {
		fmt.Fprintln(out, l)
	}
	fmt.Fprint(out, finalWorld)
}

func main() {
	Countdown(os.Stdout)
}

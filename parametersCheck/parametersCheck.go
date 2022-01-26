package main

import "fmt"

// anyPositiveWeak indicates if any value is greater than zero.
func anyPositiveWeak(values ...int) bool {
	for _, v := range values {
		if v > 0 {
			return true
		}
	}
	return false
}

// anyPositiveStrong indicates if any value is greater than zero.
// anyPositiveStrong cannot be called with less than one argument.
func anyPositiveStrong(first int, rest ...int) bool {
	if first > 0 {
		return true
	}
	for _, v := range rest {
		if v > 0 {
			return true
		}
	}
	return false
}

func main() {
	if anyPositiveWeak() {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}

	if anyPositiveStrong(1) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}

package main

func main() {
	println(sum(100))
}

//go:nosplit
func sum(n int) int {
	if n > 0 {
		return n + sum(n-1)
	} else {
		return 0
	}
}

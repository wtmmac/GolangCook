package main

func closeOfNilChannel() {
	var chan2 chan string
	close(chan2)
}

func main() {
	closeOfNilChannel()
}

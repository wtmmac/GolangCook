// task_1_1.go
package main

func sendOnClosedChannel() {
	chan1 := make(chan string)
	close(chan1)
	chan1 <- "hello"
}

func main() {
	sendOnClosedChannel()
}

package main

// Define a function type
type HandlerFunc func(str string)

// Process calls f(str).
func (f HandlerFunc) Process(str string) {
	f(str)
}

// Define testFunc with the appropriate signature
func testFunc(str string) {
	println(str)
}

func main() {
	// Convert testFunc to type of HandlerFunc
	foo := HandlerFunc(testFunc)
	foo.Process("Process..")
}

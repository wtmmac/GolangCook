package main

import (
	"os"
)

func main() {
	fd, _ := os.OpenFile("test.txt", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0670)
	print(fd)
}

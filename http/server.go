package main

import (
	"io"
	"log"
	"net/http"
	"time"
)

func HelloServer(w http.ResponseWriter, req *http.Request) {
	time.Sleep(time.Second * 10)
	io.WriteString(w, "hello, world!\n")
}

func main() {
	http.HandleFunc("/hello", HelloServer)
	err := http.ListenAndServe(":9001", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

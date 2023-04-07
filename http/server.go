package main

import (
	"log"
	"net/http"
	"time"
)

func HelloServer(wr http.ResponseWriter, r *http.Request) {
	time.Sleep(time.Second * 10)
	_, _ = wr.Write([]byte("hello, world!\n"))
}

func main() {
	http.HandleFunc("/hello", HelloServer)
	err := http.ListenAndServe(":9001", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

package main

import (
	"github.com/wtmmac/go.net/websocket"
	"net/http"
)

// EchoServer Echo the data received on the WebSocket.
func EchoServer(ws *websocket.Conn) {
	var err error
	for {
		if err = websocket.Message.Send(ws, "asdf"); err != nil {
			print("Can't send")
			break
		}
	}
}

// This example demonstrates a trivial echo server.
func main() {
	http.Handle("/echo", websocket.Handler(EchoServer))
	err := http.ListenAndServe(":12345", nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}

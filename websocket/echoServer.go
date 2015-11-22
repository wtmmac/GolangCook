package main

import (
	//"io"
	"net/http"

	"github.com/wtmmac/go.net/websocket"
)

// Echo the data received on the WebSocket.
func EchoServer(ws *websocket.Conn) {
	//io.Copy(ws, ws)
	var err error
	for {
		//io.WriteString(ws, "asdf")
		if err = websocket.Message.Send(ws, "asdf"); err != nil {
			print("Can't send")
			break
		}
	}

	//	_, err := ws.Write([]byte("echo pong"))
	//	if err != nil {
	//		panic("EchoServer: " + err.Error())
	//	}
}

// This example demonstrates a trivial echo server.
func main() {
	http.Handle("/echo", websocket.Handler(EchoServer))
	err := http.ListenAndServe(":12345", nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}

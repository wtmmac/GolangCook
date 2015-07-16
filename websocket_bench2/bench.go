package main

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/wtmmac/go.net/websocket"
	"log"
	"net/url"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// tester
func websocketTester(URL string, gorutineNumber int) {
	parsedURL, err := url.Parse(URL)

	if err != nil {
		fmt.Println(err)
	}

	ws, err := websocket.Dial(URL, "", "http://"+parsedURL.Host)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf(">>>websocketTester started:%7d\n", gorutineNumber)

	message := []byte("hello, world!")
	_, err = ws.Write(message)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Sended: %s\n\n", message)

	var dataBuffer bytes.Buffer
	var n int
	var msg = make([]byte, 512)
	for {
		n, err = ws.Read(msg)
		if err != nil {
			log.Fatal(err)
		}
		dataBuffer.Write(msg[:n])
		fmt.Printf("websocketTester %7d received\n", gorutineNumber)
		dataBuffer.Reset()
	}
}

func main() {
	var ws string
	flag.StringVar(&ws, "ws", "ws://localhost:7001/dmlive/new-msg/socket/658", "websocket address")
	var n int
	flag.IntVar(&n, "n", 1, "Number of requests to perform")
	flag.Parse()
	fmt.Printf(".\n testing target is %s\n", ws)
	fmt.Println("\x1b[43;31m\nControl + C to terminate the programme!\x1b[0m\n")

	time.Sleep(time.Second * 3)
	for i := 0; i < n; i++ {
		time.Sleep(time.Millisecond * 5)
		go websocketTester(ws, i+1)
	}

	sc := make(chan os.Signal)
	var sig os.Signal
	signal.Notify(sc, syscall.SIGINT)

	for {
		sig = <-sc
		switch sig {
		case syscall.SIGINT:
			println("SIGINT is received")
			return
		}
	}
}

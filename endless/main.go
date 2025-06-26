package main

import (
	"net/http"

	"github.com/fvbock/endless"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, world!"))
	})

	// 使用endless库启动服务
	server := endless.NewServer(":8080", mux)
	server.ListenAndServe()
}

// 重启时，只需发送USR1信号给进程即可

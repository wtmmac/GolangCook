package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	timeStart := time.Now()
	_, err := fmt.Fprintf(w, "Hello World!")
	if err != nil {
		log.Panic(err)
	}
	timeElapse := time.Since(timeStart)
	log.Println("HelloWorld:", timeElapse)
}

func timeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		timeStart := time.Now()
		//传递参数给下一个http.Handler
		next.ServeHTTP(w, r)
		timeElapse := time.Since(timeStart)
		log.Println(timeElapse)
	})
}

//新增一个中间件用于log的生成
func logMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		file, _ := os.OpenFile("test.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
		logger := log.New(file, "", log.Llongfile|log.LstdFlags)
		//将w,r传递给下一个handler
		h.ServeHTTP(w, r)
		logger.Println("请求处理完成")
	})
}
func main() {
	http.Handle("/", logMiddleware(timeMiddleware(http.HandlerFunc(HelloWorld))))
	if err := http.ListenAndServe(":9001", nil); err != nil {
		log.Panic(err)
	}
}

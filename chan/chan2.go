package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL, syscall.SIGHUP, syscall.SIGQUIT)
	<-quit
	time.Sleep(time.Second * 10)
}

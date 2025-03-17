package main

import (
	"context"
	"fmt"
	"log"
	"runtime/debug"
)

func main() {
	defer func() { //必须要先声明defer，否则不能捕获到panic异常
		fmt.Println("enter defer func")
		if err := recover(); err != nil {
			fmt.Printf("Panic occurred: %v\n", err)
			fmt.Printf("Stack trace:\n%s\n", debug.Stack())
		}
		fmt.Println("end of defer")
	}()
	raisePanic()
}

func raisePanic() {
	fmt.Println("before panic")
	// panic(errors.New("!!fatal error!"))
	raseError()
	fmt.Println("after panic")
}

func raseError() {
	InfoContextf(context.Background(), "hello")
}

func InfoContextf(ctx context.Context, format string, args ...interface{}) {
	traceId := ctx.Value("TraceKey").(string)
	prefix := fmt.Sprintf("INFO  %s ", traceId)
	_ = log.Output(2, fmt.Sprintf(prefix+format+"\n", args...))
}

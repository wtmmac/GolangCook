package main

import (
	"context"
	"fmt"
)

func test1() {
	fmt.Println("++test1++++++++++++++++++++=")
	ctx := context.Background()
	var keyA string = "keyA"
	ctxA := context.WithValue(ctx, keyA, "VALUE A")

	var keyC string = "keyA"
	ctxC := context.WithValue(ctx, keyC, "eggo")

	fmt.Println(ctxA.Value(keyA))
	//fmt.Println(ctxA.Value("keyC"))
	fmt.Println("======")
	fmt.Println(ctxC.Value(keyA))
	fmt.Println(ctxC.Value(keyC))
	fmt.Println(ctxA.Value(keyA))
	//fmt.Println(ctxC.Value("keyC"))
}

// 确认子context的key是否覆盖父context
func test2() {
	fmt.Println("\n++test2++++++++++++++++++++=")
	ctx := context.Background()
	var keyA string = "key"
	ctxA := context.WithValue(ctx, keyA, "ctxA")
	var keyC string = "key"
	ctxC := context.WithValue(ctxA, keyC, "ctxC")
	fmt.Printf("ctxA:%s  \t ctxC:%s \n", ctxA.Value(keyA), ctxC.Value(keyC))
	//ctxA:ctxA        ctxC:ctxC
	//fmt.Println(ctxC.Value(keyA)) // return empty
}

func test3() {
	fmt.Println("\n++test3++++++++++++++++++++=")
	type favContextKey string
	f := func(ctx context.Context, k favContextKey) {
		if v := ctx.Value(k); v != nil {
			fmt.Println("found value:", v)
			return
		}
		fmt.Println("key not found:", k)
	}
	k := favContextKey("language")
	ctx := context.WithValue(context.Background(), k, "Go")
	f(ctx, k)
	f(ctx, favContextKey("color"))
	// Output:
	// found value: Go
	// key not found: color
}

func main() {
	test1()
	test2()
	test3()
}

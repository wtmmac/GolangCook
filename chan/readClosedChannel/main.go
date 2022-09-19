package main

import "fmt"

func main()  {
	ch := make(chan int)
	go func() {
		ch <-1
		close(ch)
	}()


	for value := range ch{
		fmt.Println("value:", value)
	}
	//ch = nil
	fmt.Println(<- ch)
	a:=<-ch
	fmt.Println(a)

}

package main

import (
	"fmt"
	"log"
	"unsafe"
)

func MakeMem() []int {
	// mem := make([]int, 100*1024*1024)
	mem := make([]int, 1)
	// mem = append(mem, 1)
	return mem
}

func main() {
	log.Printf("run...")
	fmt.Println("Size of [100*1024*1024]int:", unsafe.Sizeof([100 * 1024 * 1024]int{})) //4000
	var a int
	fmt.Println(unsafe.Sizeof(a))
	//mem := MakeMem()
	//fmt.Println(binary.Size(mem))
	//defer log.Printf("stop.")
	//i := 0
	//c := make(chan int, 1)
	//for {
	//	go func(i int) {
	//		mem := make([]int, 100*1024*1024)
	//		log.Printf("i=%d,mem:%p", i, mem)
	//		mem[0] = <-c
	//	}(i)
	//	i++
	//	time.Sleep(200 * time.Microsecond)
	//}
}

package main

import (
	"fmt"
)

type video struct {
	name string
	iid  int
}

func (v *video) Play() {
	fmt.Println(v.name + "is playing")
}

func main() {
	video1 := video{name: "冰雪奇缘", iid: 121212}
	video1.Play()
}

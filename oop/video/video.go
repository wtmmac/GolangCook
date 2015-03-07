package video

import (
	"fmt"
)

type Video struct {
	Name string
	Iid  int
}

func (v *Video) Play() { //Public
	fmt.Println(v.Name + "is playing")
	v.sendLog()
}

func (v *Video) sendLog() { //Private
	fmt.Println("##log:" + v.Name + " is playing")
}

type VideoExt struct {
	Video   // 匿名字段
	Formats []string
}

func (v *VideoExt) ListFormat() {
	for _, format := range v.Formats {
		fmt.Println(format)
	}
}

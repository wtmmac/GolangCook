package main

import (
	"./video"
)

func main() {
	video1 := video.Video{Name: "冰雪奇缘", Iid: 121212}
	video1.Play()
	//video1.sendLog()
}

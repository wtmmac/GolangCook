package main

import (
	"github.com/wtmmac/GolangCook/oop/video"
)

func main() {
	video1 := video.Video{Name: "冰雪奇缘", Iid: 121212}
	video1.Play()
	//video1.sendLog()

	video2 := video.VideoExt{video.Video{Name: "名侦探狄仁杰", Iid: 121212}, []string{"flv", "mp4"}}
	video2.Play()
	video2.ListFormat()
}

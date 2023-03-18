package main

import (
	"encoding/json"
	"fmt"
	simplejson "github.com/bitly/go-simplejson"
)

type MediaInfo struct {
	VideoInfo interface{} `json:"video_info"`
	Title     string      `json:"Title"`
}

func main() {
	data := []byte(`{
  "person": {
    "name": {
      "first": "Leonid",
      "last": "Bugaev",
      "fullName": "Leonid Bugaev"
    },
    "github": {
      "handle": "buger",
      "followers": 109
    },
    "avatars": [
      { "url": "https://avatars1.githubusercontent.com/u/14009?v=3&s=460", "type": "thumbnail" }
    ]
  },
  "company": {
    "name": "Acme"
  }
}`)

	mediaInfo := MediaInfo{}

	mediaInfo.Title = "mediaTitle_1"
	//mediaInfo.VideoInfo = string(data)

	js, err1 := simplejson.NewJson(data) // 反序列化
	if err1 != nil {
		panic(err1.Error())
	}
	mediaInfo.VideoInfo = js

	jsonData, _ := json.Marshal(mediaInfo)
	fmt.Println(string(jsonData))

}

package main

import (
	"encoding/json"
	"fmt"
	"github.com/buger/jsonparser"
)

type MediaInfo struct {
	VideoInfo interface{} `json:"video_info"`
	Title     string      `json:"title"`
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
	// You can specify key path by providing arguments to Get function
	person, dataType, offset, err := jsonparser.Get(data, "person")
	fmt.Println(string(person))
	fmt.Println(dataType)
	fmt.Println(offset)
	fmt.Println(err)

	fmt.Println("==")

	mediaInfo.Title = "mediaInfo_1"
	mediaInfo.VideoInfo = string(data)

	jsonData, _ := json.Marshal(mediaInfo)
	fmt.Println(string(jsonData))

}

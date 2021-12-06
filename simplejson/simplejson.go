package main

import (
	"fmt"
	"github.com/bitly/go-simplejson"
)

func main() {

	body := `{"user_code":"102852181","userPic":"http://static.youku.com/user/img/avatar/80/14.jpg","pos":3,"color":16777215,"commit_time":1448265594848,"effect":0,"data":"接口就","alpha":1,"size":1,"msgtype":2,"userName":"tony大天"}`

	fmt.Println(body)
	js, err := simplejson.NewJson([]byte(body))
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(js)
	//fmt.Println(js.Get("userName").String())

	msgtype := 999
	msgtype = js.Get("msgtype").MustInt()
	fmt.Println(msgtype)
	//	fmt.Println(js.response)
}

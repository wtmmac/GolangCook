package main

import (
	//"fmt"
	"github.com/bitly/go-simplejson"
	"testing"
)

//func TestAdd(t *testing.T) {
//	r := 2
//	if r != 2 {
//		t.Errorf("not equal.")
//	}
//}

func BenchmarkJson(b *testing.B) {
	body := `{"user_code":"102852181","userPic":"http://static.youku.com/user/img/avatar/80/14.jpg","pos":3,"color":16777215,"commit_time":1448265594848,"effect":0,"data":"asdfasdfasfafsf","alpha":1,"size":1,"msgtype":2,"userName":"tony大天"}`

	for i := 0; i < b.N; i++ {
		json, err := simplejson.NewJson([]byte(body))
		if err != nil {
			panic(err.Error())
		}
		if json != nil {
			json.Get("msgtype").MustInt()
			json.Get("msgtype").MustInt()
			json.Get("msgtype").MustInt()
			json.Get("msgtype").MustInt()
			json.Get("msgtype").MustInt()
			json.Get("userName").MustString()
			json.Get("userName").MustString()
			json.Get("userName").MustString()
			json.Get("userName").MustString()
			json.Get("userName").MustString()
			json.Get("userName").MustString()
		}
		//fmt.Println(json.Get("msgtype").MustInt())
	}
	//	fmt.Println(js)
	//	fmt.Println(js.Get("response").Get("status").String())
}

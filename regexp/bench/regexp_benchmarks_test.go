package main

import (
	"regexp"
	"testing"
)

func BenchmarkRegexp(b *testing.B) {
	dm := `{"user_code":"102852181","userPic":"http://static.youku.com/user/img/avatar/80/14.jpg","pos":3,"color":16777215,"commit_time":1448265594848,"effect":0,"data":"asdfasdfasfafsf","alpha":1,"size":1,"msgtype":2,"userName":"tony大天"}`

	for i := 0; i < b.N; i++ {
		// 转义data中的数据
		reg := regexp.MustCompile(`(data":")([\S\s]+?)(",|"})`)
		findStr := reg.FindSubmatch([]byte(dm))
		dataSlashes := string(findStr[1]) + string(findStr[2]) + string(findStr[3])
		dm = reg.ReplaceAllString(dm, dataSlashes)
	}
}

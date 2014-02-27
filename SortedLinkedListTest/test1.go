package main

import (
	"GolangTest/SortedLinkedList"
	"GolangTest/danmuCache2"
	"fmt"
)

var (
	dmCache *danmuCache2.Cache
)

type WordCount struct {
	Word  string
	Count int
}

func compareValue(old, new interface{}) bool {
	if new.(WordCount).Count > old.(WordCount).Count {
		return true
	}
	return false
}

func main() {
	dmCache = danmuCache2.New(0)

	wordCounts := []WordCount{
		WordCount{"kate", 87},
		WordCount{"herry", 92},
		WordCount{"james", 81}}

	var aSortedLinkedList = SortedLinkedList.NewSortedLinkedList(10, compareValue)

	for _, wordCount := range wordCounts {
		aSortedLinkedList.PutOnTop(wordCount)
	}
	for element := aSortedLinkedList.List.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value.(WordCount))
	}

	// set cache
	id := "1"
	danmu := make(danmuCache2.Danmu)

	// 第一分钟
	danmu["1"] = []danmuCache2.DanmuMinute{
		danmuCache2.DanmuMinute{1, "测试", 1300002101},
		danmuCache2.DanmuMinute{2, "测试", 1300002101},
		danmuCache2.DanmuMinute{3, "测试", 1300002101}}

	// 第二分钟
	danmu["2"] = []danmuCache2.DanmuMinute{
		danmuCache2.DanmuMinute{0, "测试", 1300002101},
		danmuCache2.DanmuMinute{0, "测试", 1300002101},
		danmuCache2.DanmuMinute{0, "测试", 1300002101}}

	dmCache.Add(id, danmu)

	// get cache
	if v, ok := dmCache.Get("1"); ok {
		// 返回弹幕
		fmt.Println(v.String())
		fmt.Println("==\t==")
		fmt.Println(v["1"][1].Content)
	} else {
		fmt.Println("no value")
	}
}

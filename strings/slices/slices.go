package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type CommData struct {
	Uid   string
	Ctime int
	Md5   string
}

type CommSlice []CommData

func (p CommSlice) Len() int           { return len(p) }
func (p CommSlice) Less(i, j int) bool { return p[i].Ctime < p[j].Ctime }
func (p CommSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func main() {
	data := []string{"121-1237-asdf", "121-1234-asdf", "122-1235-kakkaas", "121-1236-asdf"}
	comms := CommSlice{}
	for _, val := range data {
		t := strings.Split(val, "-")
		ctime, _ := strconv.Atoi(t[1])
		comms = append(comms, CommData{
			Uid:   t[0],
			Ctime: ctime,
			Md5:   t[2],
		})
	}
	sort.Sort(comms)
	fmt.Println(comms)
	size := len(comms)
	for i := 0; i < size; i++ {
		fmt.Println(comms[size-1-i])
	}
}

package main

import (
	"encoding/json"
	"sync"
)


type Student struct {
	Name   string
	Age    int32
	Remark [1024]byte
}

var buf, _ = json.Marshal(Student{Name: "Geektutu", Age: 25})

func unmarsh() {
	stu := &Student{}
	json.Unmarshal(buf, stu)
}

var studentPool = sync.Pool{
	New: func() interface{} {
		return new(Student)
	},
}

func main() {
	stu := studentPool.Get().(*Student)
	json.Unmarshal(buf, stu)
	studentPool.Put(stu)
}

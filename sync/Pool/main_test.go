package main

import (
	"bytes"
	"encoding/json"
	"sync"
	"testing"
)

// go test -bench . -benchmem

/**
在这个例子中，因为 Student 结构体内存占用较小，内存分配几乎不耗时间。
而标准库 json 反序列化时利用了反射，效率是比较低的，占据了大部分时间，
因此两种方式最终的执行时间几乎没什么变化。但是内存占用差了一个数量级，
使用了 sync.Pool 后，内存占用仅为未使用的 234/5096 = 1/22，对 GC 的影响就很大了。
*/
func BenchmarkUnmarshal(b *testing.B) {
	for n := 0; n < b.N; n++ {
		stu := &Student{}
		json.Unmarshal(buf, stu)
	}
}

func BenchmarkUnmarshalWithPool(b *testing.B) {
	for n := 0; n < b.N; n++ {
		stu := studentPool.Get().(*Student)
		json.Unmarshal(buf, stu)
		studentPool.Put(stu)
	}
}

/**
这个例子创建了一个 bytes.Buffer 对象池，而且每次只执行一个简单的 Write 操作，存粹的内存搬运工，
耗时几乎可以忽略。而内存分配和回收的耗时占比较多，因此对程序整体的性能影响更大
*/

var bufferPool = sync.Pool{
	New: func() interface{} {
		return &bytes.Buffer{}
	},
}

var data = make([]byte, 10000)

func BenchmarkBufferWithPool(b *testing.B) {
	for n := 0; n < b.N; n++ {
		buf := bufferPool.Get().(*bytes.Buffer)
		buf.Write(data)
		buf.Reset()
		bufferPool.Put(buf)
	}
}

func BenchmarkBuffer(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var buf bytes.Buffer
		buf.Write(data)
	}
}

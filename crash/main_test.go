package main

import (
	"fmt"
	"reflect"
	"testing"
	"unsafe"
)

func BenchmarkMem(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mem := MakeMem()
		// fmt.Println(binary.Size(mem))
		fmt.Println(unsafe.Sizeof(mem))
	}
}

func TestMakeMem(t *testing.T) {
	tests := []struct {
		name string
		want []int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MakeMem(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MakeMem() = %v, want %v", got, tt.want)
			}
		})
	}
}

package main

import (
	"encoding/binary"
	"fmt"
	"reflect"
	"testing"
)

func BenchmarkMem(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mem := MakeMem()
		fmt.Println(binary.Size(mem))
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

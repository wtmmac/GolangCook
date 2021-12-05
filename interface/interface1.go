package main

import (
	"fmt"
	"math"
)

var string1 interface{}

type Shape2D interface {
	Perimeter() float64
}
type circle struct {
	R float64
}
func (c circle) Perimeter() float64 {
	return 2 * math.Pi * c.R
}

func main() {
	string1 = "asdf"
	fmt.Println(string1)

	if k, ok := string1.(string); ok {
		fmt.Println("k is string:", k)
	}

	a := circle{R: 1.5}
	fmt.Printf("R %.2f -> Perimeter %.3f \n", a.R, a.Perimeter())
	_, ok := interface{}(a).(Shape2D)
	if ok {
		fmt.Println("a is a Shape2D!")
	}
}

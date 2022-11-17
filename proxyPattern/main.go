package main

import "fmt"

type Subject interface {
	Do() string
}

type RelSubject struct {
}

func (r *RelSubject) Do() string {
	return "test"
}

type Proxy struct {
	real RelSubject
}

func (proxy *Proxy) Do() string {
	res := proxy.real.Do()
	return res
}

func main() {
	var sub Subject
	sub = &Proxy{}

	res := sub.Do()

	fmt.Println(res)
}

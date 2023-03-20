package main

import "fmt"

func main() {
	d := []byte{'a', 'b', 'a', 'd'}
	e := d[0:2]

	fmt.Println(e)

	keys := make([]byte, 0, len(d))

	fmt.Println(len(keys))
}

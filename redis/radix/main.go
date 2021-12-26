package main

import (
	"fmt"
	"github.com/mediocregopher/radix/v3"
)

func main() {
	client, err := radix.NewPool("tcp", "127.0.0.1:6379", 10)
	if err != nil {
		// handle error
	}

	err = client.Do(radix.Cmd(nil, "SET", "foo", "someval"))
	if err != nil {
		// handle error
	}

	var fooVal string
	err = client.Do(radix.Cmd(&fooVal, "GET", "foo"))
	if err != nil {
		// handle error
	}

	fmt.Println(fooVal)

	var fooValB []byte
	err = client.Do(radix.Cmd(&fooValB, "GET", "foo"))

}

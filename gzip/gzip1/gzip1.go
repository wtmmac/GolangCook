package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
)

func main() {
	str := `gzip test string`
	var buf bytes.Buffer
	var z = gzip.NewWriter(&buf)
	_, err := z.Write([]byte(str))
	if err != nil {
		panic(err)
	}

	err = z.Close()
	if err != nil {
		panic(err)
	}
	fmt.Println("gzip:", buf.String())
}

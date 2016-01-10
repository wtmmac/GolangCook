package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func findLastId() int {
	filename := "exampleData.txt"
	info, _ := os.Stat(filename)
	filesize := info.Size()

	var data string
	id := 0

	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	buf := make([]byte, 1)

	findLineFeed := 0

	for i := 1; i <= 1024; i++ {
		offset := filesize - int64(i)
		_, err := f.ReadAt(buf, offset)
		if err != nil || io.EOF == err {
			panic(err)
		}
		// 找到倒数第二个换行符
		if string(buf) == "\n" {
			findLineFeed++
			if findLineFeed == 2 {
				offset = offset + 1
				f.Seek(offset, 0)
				rd := bufio.NewReader(f)
				data, err = rd.ReadString('\n')

				if err != nil {
					fmt.Println(err)
				}

				dataSlice := strings.Split(data, "|")
				if len(dataSlice) > 0 {
					id, _ = strconv.Atoi(dataSlice[1])
				}
				//fmt.Println(data)
				break
			}
		}
	}
	return id
}

func main() {
	id := findLastId()
	fmt.Println(id)
}

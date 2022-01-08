package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("starting..")
	logFile := "log.txt"
	f, err := os.Open(logFile)
	defer f.Close()
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}

	//将文件作为一个io.Reader对象进行buffered I/O操作
	br := bufio.NewReader(f)
	for {
		//每次读取一行
		line, err := br.ReadString('\n')
		if err == io.EOF {
			break
		} else {
			//	fmt.Printf("%v", line)
			task := strings.Split(strings.TrimRight(line, "\r\n"), " ")
			data, err := strconv.Atoi(task[0])
			if err != nil {
				fmt.Printf("%v\n", err)
				os.Exit(1)
			}
			//Charging Standard 0.36/G
			result := float32(data) * 0.0005 * 0.36
			fmt.Printf("%s,%6.2f\n", task[1], result)
		}
	}

}

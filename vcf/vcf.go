package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	//"reflect"
	"strings"
)

type Vcard struct {
	BEGIN   string
	VERSION string
	N       string
	FN      string
	UID     string
	TEL     string
	END     string
}

func main() {
	vcfFile := "test.vcf"
	f, err := os.Open(vcfFile)
	defer f.Close()
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}

	//var tmpVcard Vcard
	var content_lines = make([]string, 10)

	//将文件作为一个io.Reader对象进行buffered I/O操作
	br := bufio.NewReader(f)
	for {
		//每次读取一行
		line, err := br.ReadString('\n')
		if err == io.EOF {
			break
		} else {
			//fmt.Printf("%v", line)
			content_lines = append(content_lines, strings.TrimRight(line, "\r\n"))
		}
	}

	//var typeof reflect.Type = reflect.TypeOf(tmpVcard)
	//for item := 0; item < typeof.NumField(); item++ {
	//	f := typeof.Field(item)
	//	fmt.Println(f.Name)
	//}

	for k, v := range content_lines {
		if strings.HasPrefix(v, "N") {
			tmpName := strings.Split(content_lines[k+1], ":")[1]

			r := []rune(tmpName)
			fmt.Print("N:", string(r[0]), ";")
			for i := 1; i < len(r); i++ {
				fmt.Print(string(r[i]))
			}
			fmt.Print(";;;\r\n")
		} else {
			fmt.Println(v)
		}

	}
}

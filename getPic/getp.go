package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
)

var urlList []string
var album chan string
var dir string

func fillUrlList() {
	f, err := os.Open("urllist.txt")
	defer f.Close()
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}

	// 将文件作为一个io.Reader对象进行buffered I/O操作
	br := bufio.NewReader(f)
	for {
		// 每次读取一行
		line, err := br.ReadString('\n')
		if err == io.EOF {
			break
		} else {
			urlList = append(urlList, strings.ToLower(strings.TrimRight(line, "\r\n")))
		}
	}
}

func main() {
	defer func() { //必须要先声明defer，否则不能捕获到panic异常
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	dir = "tmp/"
	err := os.Mkdir(dir, 0777)
	if err != nil {
		isexist := os.IsExist(err)
		log.Println("dir exist is :", isexist)
	}
	album = make(chan string, 200)
	fillUrlList()
	for _, v := range urlList {
		GetAlbum(v)
	}

}

func GetAlbum(url string) {
	data := GetUrl(url)
	body := string(data)
	part := regexp.MustCompile(`photolst_photo" title="(.*)">\s*<img src="(.*)"`)
	match := part.FindAllStringSubmatch(body, -1)
	for _, v := range match {
		album <- v[2]
		go GetImage()
	}
}

func GetImage() {
	imgUrl := <-album
	str := strings.Split(imgUrl, "/")
	length := len(str)
	source := GetUrl(imgUrl)
	name := str[length-1]
	file, err := os.Create(dir + name)
	if err != nil {
		panic(err)
	}
	size, err := file.Write(source)
	defer file.Close()
	if err != nil {
		panic(err)
	}
	log.Println(size)
}

func GetUrl(url string) []byte {
	ret, err := http.Get(url)
	if err != nil {
		log.Println(url)
		status := map[string]string{}
		status["status"] = "400"
		status["url"] = url
		panic(status)
	}
	body := ret.Body
	data, _ := ioutil.ReadAll(body)
	return data
}

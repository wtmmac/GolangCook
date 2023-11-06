package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	urlIndex := "http://www.example.com/page/"
	for i := 1; i <= 4; i++ {
		fetchList(urlIndex + strconv.Itoa(i))
	}
}

func fetchList(url string) {
	dir, _ := os.Getwd()
	fmt.Println(dir)

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	dom, _ := goquery.NewDocumentFromReader(res.Body)
	dom.Find(".post-title-link").Each(func(i int, selection *goquery.Selection) {
		href, _ := selection.Attr("href")
		fetchContent(href)
	})
}

func WriteFile(filename string, content []byte) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = file.Write(content)
	if err != nil {
		panic(err)
	}

	fmt.Println("File created and content written successfully.")
}

func fetchContent(url string) {
	dir, _ := os.Getwd()
	fmt.Println(dir)

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	dom, _ := goquery.NewDocumentFromReader(res.Body)
	dom.Find(".post-block").Each(func(i int, selection *goquery.Selection) {
		postTitle := strings.TrimSpace(selection.Find(".post-title").Text())
		filename := postTitle + ".html"

		fmt.Println(selection.Find(".post-body").Html())
		postBody, _ := selection.Find(".post-body").Html()
		postBody = "<H1>" + postTitle + "</H1>" + postBody

		WriteFile(filename, []byte(postBody))

		selection.Find("img").Each(func(j int, selection *goquery.Selection) {
			imgSrc, _ := selection.Attr("src")
			download(imgSrc)
		})
	})
}

func download(url string) {
	client := http.Client{}
	response, err := client.Get(url)
	if err != nil {
		fmt.Println("获取失败")
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		fmt.Println("响应码错误")
	}

	urls := strings.Split(url, "/")
	filename := urls[len(urls)-1]

	data, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("读取响应数据失败")
	}
	err = os.WriteFile(filename, data, 0666)
	if err != nil {
		fmt.Println("写入文件失败")
	}
}

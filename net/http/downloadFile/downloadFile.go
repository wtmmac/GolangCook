package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	url := "https://img-blog.csdnimg.cn/img_convert/7d12ff8917f2ffba51e6c8481a5ff7e7.png"
	download(url)
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

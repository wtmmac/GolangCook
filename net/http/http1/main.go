// 基本的GET请求
package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	//_ = trace.Start(os.Stderr)
	//defer trace.Stop()
	resp, err := http.Get("http://httpbin.org/get")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	fmt.Println(string(body))
	fmt.Println(resp.StatusCode)

	if resp.StatusCode == 200 {
		fmt.Println("ok")
	}

}

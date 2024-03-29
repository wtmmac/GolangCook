package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func httpGet(http_url string) string {
	if !strings.HasPrefix(http_url, "http") {
		http_url = "http://" + http_url
	}

	// validate url
	host, err := url.ParseRequestURI(http_url)
	if err != nil {
		fmt.Println(err)
	}

	// validate host
	_, err = net.LookupIP(host.Host)
	if err != nil {
		fmt.Println(err)
	}

	client := &http.Client{
		Transport: &http.Transport{
			Dial: func(network, address string) (net.Conn, error) {
				deadline := time.Now().Add(10 * time.Second)
				c, err := net.DialTimeout(network, address, 5*time.Second)
				if err != nil {
					return nil, err
				}

				c.SetDeadline(deadline)
				return c, nil
			},
		},
	}

	req, err := http.NewRequest("GET", http_url, nil)
	if err != nil {
		fmt.Println(err)
	}

	// gzip
	req.Header.Set(
		"User-Agent",
		"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.11 (KHTML, like Gecko) Chrome/23.0.1271.64 Safari/537.11",
	)
	req.Header.Set("Accept-Encoding", "gzip")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	var reader io.ReadCloser
	switch resp.Header.Get("Content-Encoding") {
	case "gzip":
		reader, _ = gzip.NewReader(resp.Body)
		fmt.Println("With gzip")
		defer reader.Close()
	default:
		reader = resp.Body
		fmt.Println("No gzip")
	}
	// fmt.Println(reader)
	body, _ := io.ReadAll(reader)
	return string(body)
}

func main() {
	result := httpGet("http://www.baidu.com")
	fmt.Println(result)

	// var result2 = httpGet("https://auth.alipay.com/login/index.htm")
	// fmt.Println(result2)
}

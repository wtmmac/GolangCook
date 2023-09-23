package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	url := "http://live.comments.tudou.com/dispatcher2/dm/add/?dm=%7B%22effect%22%3A0%2C%22alpha%22%3A1%2C%22size%22%3A1%2C%22pos%22%3A3%2C%22color%22%3A16711680%2C%22commit_time%22%3A1444465369879%2C%22user_code%22%3A%2246313317%22%2C%22data%22%3A%22%E6%8E%A5%E5%8F%A3%E5%B0%B1%22%7D&liveid=14"

	req, _ := http.NewRequest("GET", url, nil)

	// req.Header.Add("referer", "http//www.tudou.com")
	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("postman-token", "b5b74697-d0e0-13ab-6ef1-96f786a7af4a")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}

package main

import (
	"fmt"
	"github.com/bitly/go-simplejson"
	//	"github.com/liudng/godump"
)

func main() {

	body := `
  {"response": {
  "status": "SUCCESS",
  "data": {
    "mxRecords": [
      {
        "value": "us2.mx3.mailhostbox.com.",
        "ttl": 1,
        "priority": 100,
        "hostName": "hostname1"
      }
    ]
	}}}`
	js, err := simplejson.NewJson([]byte(body))
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(js)
	fmt.Println(js.Get("response").Get("status").String())
	//	fmt.Println(js.response)
}

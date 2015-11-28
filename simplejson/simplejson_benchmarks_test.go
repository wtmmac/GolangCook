package main

import (
	//	"fmt"
	"github.com/bitly/go-simplejson"
	"testing"
)

//func TestAdd(t *testing.T) {
//	r := 2
//	if r != 2 {
//		t.Errorf("not equal.")
//	}
//}

func BenchmarkJson(b *testing.B) {
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

	for i := 0; i < b.N; i++ {
		_, err := simplejson.NewJson([]byte(body))
		if err != nil {
			panic(err.Error())
		}
	}
	//	fmt.Println(js)
	//	fmt.Println(js.Get("response").Get("status").String())
}

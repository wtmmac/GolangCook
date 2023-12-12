package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"math"
	"net/http"
	"os"
	"strconv"

	"github.com/bitly/go-simplejson"
)

const pageSize = 1
const url = "http://10.19.15.99:9200/video_push_sender_consumer_log_prod-2023.12.10/_search"

func main() {
	json := esQueryDSL(0)
	total, err := parseTotalRecords(json)
	if err != nil {
		panic(err)
	}
	fmt.Printf("total: %d\n", total)
	page := int(math.Ceil(float64(total) / float64(pageSize)))
	fmt.Printf("page: %d\n", page)

	for i := 0; i <= page; i++ {
		json := esQueryDSL(i)
		fmt.Println(json)
		// if i > 5 {
		// 	break
		// }
	}
}

func esQueryDSL(page int) *simplejson.Json {
	fmt.Printf("query page: %d\n", page)
	queryDsl := fmt.Sprintf(`{
  "from": ` + strconv.Itoa(page*pageSize) + `, 
  "size": ` + strconv.Itoa(pageSize) + `, 
   "query": {
    "bool": {
      "must": [
        {
          "query_string": {
            "query": "",
            "analyze_wildcard": true,
            "default_field": "*"
          }
        },
         {
          "range": {
            "@timestamp": {
              "gte": 1687104000000,
              "lte": 1687190399999,
              "format": "epoch_millis"
            }
          }
        }
      ],
      "filter": [],
      "should": [],
      "must_not": []
    }
  },
  "_source": ["msg"]
}`)

	var requestBody = []byte(queryDsl)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
		panic(err)
	}
	if len(os.Getenv("esUsername")) == 0 {
		panic("need esUsername")
	}
	req.SetBasicAuth(os.Getenv("esUsername"), os.Getenv("esPassword"))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err.Error())
		}
	}(resp.Body)

	body, _ := ioutil.ReadAll(resp.Body)
	json, err1 := simplejson.NewJson(body)
	if err1 != nil {
		panic(err1.Error())
	}
	return json
}

func parseTotalRecords(js *simplejson.Json) (int, error) {
	if data, ok := js.Get("hits").CheckGet("total"); ok {
		total, totalErr := data.Int()
		return total, totalErr
	}
	return 0, errors.New("js.Get error")
}

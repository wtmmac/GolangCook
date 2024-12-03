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

const pageSize = 10
const url = "http://10.18.19.11:9200/logstash-api.tv.sohu.com-2024.05.23/_search"

func main() {
	json := esQueryDSL(0)
	fmt.Println(json)
	total, err := parseTotalRecords(json)
	if err != nil {
		panic(err)
	}
	fmt.Printf("total: %d\n", total)
	page := int(math.Ceil(float64(total) / float64(pageSize)))
	fmt.Printf("page: %d\n", page)

	for i := 0; i <= page; i++ {
		// if i > 1 {
		// 	break
		// }
		json := esQueryDSL(i)
		result, _ := json.MarshalJSON()
		fmt.Println(string(result))
	}
}

func esQueryDSL(page int) *simplejson.Json {
	// fmt.Printf("query page: %d\n", page)
	queryDsl := fmt.Sprintf(`{
  "from": ` + strconv.Itoa(page*pageSize) + `, 
  "size": ` + strconv.Itoa(pageSize) + `, 
  "track_total_hits": true,
  "query": {
    "bool": {
      "must": [],
      "filter": [
        {
          "range": {
            "@timestamp": {
              "format": "strict_date_optional_time",
              "gte": "2024-05-23T13:30:00.000Z",
              "lte": "2024-05-23T14:15:00.000Z"
            }
          }
        },
        {
          "match_phrase": {
            "uri": "/mobile_user/push/uploadtoken.json"
          }
        }
      ],
      "should": [],
      "must_not": [
        {
          "match_phrase": {
            "status": "200"
          }
        }
      ]
    }
  },
  "_source": true
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
	// fmt.Println(string(body))
	// os.Exit()
	json, err1 := simplejson.NewJson(body)
	if err1 != nil {
		panic(err1.Error())
	}
	return json
}

func parseTotalRecords(js *simplejson.Json) (int, error) {
	if data, ok := js.Get("hits").Get("total").CheckGet("value"); ok {
		total, totalErr := data.Int()
		return total, totalErr
	}
	return 0, errors.New("js.Get error")
}

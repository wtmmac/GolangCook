package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/bitly/go-simplejson"
)

const pageSize = 10
const url = "http://10.18.19.11:9200/logstash-api.my.tv.sohu.com-2024.12.03/_search"

func rangeTime() (startTime, endTime string) {
	// 设置时区为东八区（中国标准时间）
	location, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println("Error loading location:", err)
		return
	}

	// 获取当前时间
	now := time.Now()
	beforeOneHour := now.Add(time.Duration(-1 * time.Hour))

	nowInShanghai := now.In(location)
	endTime = nowInShanghai.Format("2006-01-02T15:04:05.000Z07:00")

	beforeOneHourInShanghai := beforeOneHour.In(location)
	startTime = beforeOneHourInShanghai.Format("2006-01-02T15:04:05.000Z07:00")

	return startTime, endTime
}
func main() {
	startTime, endTime := rangeTime()
	json := esQueryDSL(startTime, endTime)
	total, uids, err := parseTotalRecords(json)
	if err != nil {
		panic(err)
	}
	fmt.Printf("total: %d\n", total)
	fmt.Printf("uids: %d\n", uids)
}

func esQueryDSL(startTime, endTime string) *simplejson.Json {
	queryDsl := fmt.Sprintf(`{
  	"size": 0, 
  	"query": {
		"bool": {
			"filter": [
				{
					"range": {
						"@timestamp": {
						"format": "strict_date_optional_time",
						"gte": "` + startTime + `",
						"lte": "` + endTime + `"
						}
					}
				},
				{
					"match_phrase": {
						"uri": "/comment"
					}
				},
				{
					"match_phrase": {
						"status": "404"
					}
				}
			]
		}
  	},
	"aggs": {
		"uids": {
			"cardinality": {
				"field": "uid.keyword"
			}
		},
		"total": {
			"value_count": {
				"field": "_index"
			}
		},
		"agg_uri": {
			"terms": {
				"field": "uri.keyword",
				"order": {
					"1": "desc"
				},
				"size": 10
			},
			"aggs": {
				"1": {
					"cardinality": {
						"field": "uid.keyword"
					}
				},
				"2": {
					"cardinality": {
						"field": "clientIp.keyword"
					}
				}
			}
		}
	}
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

	body, _ := io.ReadAll(resp.Body)
	fmt.Println(string(body))
	json, err1 := simplejson.NewJson(body)
	if err1 != nil {
		panic(err1.Error())
	}
	return json
}

func parseTotalRecords(js *simplejson.Json) (total, uids int, err error) {
	if dataTotal, ok := js.Get("aggregations").Get("total").CheckGet("value"); ok {
		total, err = dataTotal.Int()
		if err != nil {
			return total, uids, err
		}
	}
	if uidsTotal, ok := js.Get("aggregations").Get("uids").CheckGet("value"); ok {
		uids, err = uidsTotal.Int()
		if err != nil {
			return total, uids, err
		}
	}
	return total, uids, nil
}

package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/bitly/go-simplejson"
)

const url = "http://10.18.19.11:9200/logstash-api.my.tv.sohu.com-2024.12.03,logstash-api.my.tv.sohu.com-2024.12.02/_search"

func main() {
	url := GenerateURL()
	fmt.Println("Generated URL:", url)

	startTime, endTime, err := rangeTime()
	if err != nil {
		panic(err)
	}
	json, err := esQueryDSL(startTime, endTime)
	if err != nil {
		panic(err)
	}
	total, uids, agg, err := parseTotalRecords(json)
	if err != nil {
		panic(err)
	}

	fmt.Printf("total: %d\n", total)
	fmt.Printf("uids: %d\n", uids)
	fmt.Printf("=====uri统计=====\n%s\n", agg)
}

func esQueryDSL(startTime, endTime string) (*simplejson.Json, error) {
	queryDsl := fmt.Sprintf(`{
  	"size": 20, 
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
		"uid": {
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

	// fmt.Println(queryDsl)
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
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err.Error())
		}
	}(resp.Body)

	body, _ := io.ReadAll(resp.Body)
	fmt.Println(string(body))
	json, err := simplejson.NewJson(body)
	if err != nil {
		return nil, err
	}
	return json, nil
}

func parseTotalRecords(js *simplejson.Json) (total, uids int, agg string, err error) {
	if dataTotal, ok := js.Get("aggregations").Get("total").CheckGet("value"); ok {
		total, err = dataTotal.Int()
		if err != nil {
			return total, uids, agg, err
		}
	}
	if uidsTotal, ok := js.Get("aggregations").Get("uid").CheckGet("value"); ok {
		uids, err = uidsTotal.Int()
		if err != nil {
			return total, uids, agg, err
		}
	}

	// 获取 aggregations -> agg_uri -> buckets
	buckets := js.GetPath("aggregations", "agg_uri", "buckets")
	if buckets == nil {
		return total, uids, agg, fmt.Errorf("aggregations -> agg_uri -> buckets not found")
	}

	// 构建结果字符串
	var result strings.Builder

	// 遍历 buckets 并构建每个 bucket 的信息
	for _, bucket := range buckets.MustArray() {
		bucketMap := bucket.(map[string]interface{})
		key := bucketMap["key"].(string)
		docCount := bucketMap["doc_count"]

		// 添加 key 和 doc_count 到结果字符串
		result.WriteString(fmt.Sprintf("uri: %s, Doc Count: %s\n", key, docCount))

		// 如果需要进一步处理 1 和 2 的 value
		value1 := bucketMap["1"].(map[string]interface{})["value"]
		value2 := bucketMap["2"].(map[string]interface{})["value"]
		result.WriteString(fmt.Sprintf("uid: %s, clientIp: %s\n", value1, value2))
		result.WriteString(strings.Repeat("-", 20) + "\n")
	}

	return total, uids, result.String(), nil
}

func rangeTime() (startTime, endTime string, err error) {
	location, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		return "", "", err
	}

	now := time.Now()
	beforeOneHour := now.Add(time.Duration(-24 * time.Hour))

	nowInShanghai := now.In(location)
	endTime = nowInShanghai.Format("2006-01-02T15:04:05.000Z07:00")

	beforeOneHourInShanghai := beforeOneHour.In(location)
	startTime = beforeOneHourInShanghai.Format("2006-01-02T15:04:05.000Z07:00")

	return startTime, endTime, nil
}

// GenerateURL 根据当前时间和倒推1小时的时间生成 URL
func GenerateURL() string {
	const baseURL = "http://10.18.19.11:9200/"
	const indexPrefix = "logstash-api.my.tv.sohu.com-"
	const searchPath = "/_search"

	// 获取当前时间和倒推1小时的时间
	now := time.Now()
	hourAgo := now.Add(-24 * time.Hour)

	// 格式化日期
	currentDate := now.Format("2006.01.02")
	hourAgoDate := hourAgo.Format("2006.01.02")

	// 生成索引部分
	indexPart := indexPrefix + currentDate
	if currentDate != hourAgoDate {
		indexPart += "," + indexPrefix + hourAgoDate
	}

	// 生成完整的 URL
	return baseURL + indexPart + searchPath
}

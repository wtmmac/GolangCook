package main

import (
	"bytes"
	"es/query"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/bitly/go-simplejson"
)

// Colors
const (
	Reset       = "\033[0m"
	Red         = "\033[31m"
	Green       = "\033[32m"
	Yellow      = "\033[33m"
	Blue        = "\033[34m"
	Magenta     = "\033[35m"
	Cyan        = "\033[36m"
	White       = "\033[37m"
	BlueBold    = "\033[34;1m"
	MagentaBold = "\033[35;1m"
	RedBold     = "\033[31;1m"
	YellowBold  = "\033[33;1m"
)

func main() {
	if err := query.QueryManager.ExecuteAllQueries(ExecuteQuery); err != nil {
		fmt.Println("Error executing queries:", err)
	} else {
		fmt.Println("All queries executed successfully.")
	}
}

func ExecuteQuery(name, query, index string) error {
	fmt.Printf("====== Task: %s\n", name)
	// fmt.Printf("Executing query on index %s\n", index)
	// fmt.Printf("%s\n", strings.ReplaceAll(query, "\t", "   "))

	json, err := esQueryDSL(query, index)
	if err != nil {
		return err
	}
	total, uids, agg, err := parseTotalRecords(json)
	if err != nil {
		return err
	}
	fmt.Printf(RedBold+"total: %d\n"+Reset, total)
	fmt.Printf("uids: %d\n", uids)
	fmt.Printf("------ uri statistics ------\n%s\n", agg)

	return nil
}

func esQueryDSL(queryDsl, index string) (*simplejson.Json, error) {
	var requestBody = []byte(queryDsl)
	req, err := http.NewRequest("POST", index, bytes.NewBuffer(requestBody))
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
	// fmt.Println(string(body))

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

	// aggregations -> agg_uri -> buckets
	buckets := js.GetPath("aggregations", "agg_uri", "buckets")
	if buckets == nil {
		return total, uids, agg, fmt.Errorf("aggregations -> agg_uri -> buckets not found")
	}

	var result strings.Builder

	for _, bucket := range buckets.MustArray() {
		bucketMap := bucket.(map[string]interface{})
		key := bucketMap["key"].(string)
		docCount := bucketMap["doc_count"]

		result.WriteString(fmt.Sprintf("uri: %s, Doc Count: %s\n", key, docCount))

		value1 := bucketMap["1"].(map[string]interface{})["value"]
		value2 := bucketMap["2"].(map[string]interface{})["value"]
		result.WriteString(fmt.Sprintf("uid: %s, clientIp: %s\n", value1, value2))
		result.WriteString(strings.Repeat("-", 20) + "\n")
	}

	return total, uids, result.String(), nil
}

package main

import (
	"fmt"

	"github.com/bitly/go-simplejson"
)

func main() {

	body := `{"amount":{"currency":"CNY","discount_refund":0,"from":[],"payer_refund":3,"payer_total":3,"refund":3,"settlement_refund":3,"settlement_total":3,"total":3},"channel":"ORIGINAL","create_time":"2024-07-20T16:47:24+08:00","funds_account":"AVAILABLE","out_refund_no":"RE20240720101817892490","out_trade_no":"20240720081558103679","promotion_detail":[],"refund_id":"50302900262024072070910625700","status":"PROCESSING","transaction_id":"4200002313202407205889648260","user_received_account":"招商银行借记卡4246"}`

	// fmt.Println(body)
	resp, err := simplejson.NewJson([]byte(body))
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(resp)
	stringtemp := resp.Get("transaction_id").MustString()
	fmt.Println(stringtemp + ":")

}

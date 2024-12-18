package main

import (
	"fmt"
	"time"

	"github.com/devfeel/mapper"
)

func main() {
	t := time.Now()
	fmt.Println(t.Format("15:04:05"))
	fmt.Println(t.Second())
	fmt.Println(time.Unix(1136185445, 0))
	fmt.Println(t.Local().GoString())

	var oldVipExpireTime time.Time
	newVipExpireTime := time.Now().AddDate(0, 3, 0)
	snapshot := make(map[string]interface{})
	snapshot["donate"] = 1
	snapshot["donate"] = 3
	snapshot["oldVipExpireTime"] = oldVipExpireTime.Format("2006-01-02 15:04:05")
	snapshot["newVipExpireTime"] = newVipExpireTime.Format("2006-01-02 15:04:05")
	snapshotBytes, err := mapper.MapToJson(snapshot)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(snapshotBytes))

}

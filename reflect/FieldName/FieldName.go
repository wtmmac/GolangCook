package main

import (
	"fmt"
	"reflect"
)

type Vcard struct {
	BEGIN   string
	VERSION string
	N       string
	FN      string
	UID     string
	TEL     string
	END     string
}

func main() {
	var tmpVcard Vcard = Vcard{"BEGIN", "VERSION", "N", "FN", "UID", "TEL", "END"}
	var typeof reflect.Type = reflect.TypeOf(tmpVcard)
	for item := 0; item < typeof.NumField(); item++ {
		f := typeof.Field(item)
		fmt.Println(f.Name)
	}
}

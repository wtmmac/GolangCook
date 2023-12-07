package main

import (
	"fmt"
	"reflect"
)

func Walk(x interface{}, fn func(input string)) {
	val := reflect.ValueOf(x)

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)

		switch field.Kind() {
		case reflect.String:
			fn(field.String())
		case reflect.Struct:
			fmt.Printf(("this a struct:%v"), field.Interface())
			Walk(field.Interface(), fn)
		}
	}
}

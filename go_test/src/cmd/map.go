package main

import (
	"fmt"

)

func maps() {
	ages := map[string]int {
	    "lili" : 13,
	    "nick":23,
	    "jacky":55,
	}
	
	dic := map[string]string {
		"tony":"tonywang",
		"lisa":"lisazhang",
	}
	
	fmt.Println(ages)
	fmt.Println(dic)

}

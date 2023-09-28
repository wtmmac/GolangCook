package main

import "fmt"

func main() {
	d := []byte{'a', 'b', 'a', 'd'}
	e := d[0:2]

	fmt.Println(e)

	keys := make([]byte, 0, len(d))

	fmt.Println(len(keys))

	urls := make(map[string]string, 3)
	urls["baidu"] = "www.baidu.com"
	urls["google"] = "www.google.com"
	urls["csdn"] = "www.csdn.net"

	// bad method
	names := make([]string, len(urls))
	for key := range urls {
		names = append(names, key)
	}
	fmt.Printf("%#v ,%d\n", names, len(names))

	// good method
	namesGood := make([]string, 0)
	for key := range urls {
		namesGood = append(namesGood, key)
	}
	fmt.Printf("%#v ,%d\n", namesGood, len(namesGood))
}

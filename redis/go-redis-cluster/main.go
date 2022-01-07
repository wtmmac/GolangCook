package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {
	client := redis.NewClusterClient(&redis.ClusterOptions{
		//Addrs: []string{"127.0.0.1:6379", ":7001", ":7002", ":7003", ":7004", ":7005"},
		Addrs: []string{"10.19.10.52:29135", "10.19.17.95:31376"},
		// To route commands by latency or randomly, enable one of the following.
		//RouteByLatency: true,
		//RouteRandomly: true,
		Password: "",
	})
	defer client.Close()

	pong, err := client.Ping(ctx).Result()
	fmt.Println(pong, err)

	err = client.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := client.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := client.Get(ctx, "key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exists")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}

	client.HSet(ctx, "myhash", "key1", "value1", "key2", "value2")
}

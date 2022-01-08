package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

var (
	ctx    = context.Background()
	client = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
)

func main() {
	pong, err := client.Ping(ctx).Result()
	fmt.Println(pong, err)

	strings()
	hashes()
}

func strings() {
	fmt.Println("\nstrings started.======")
	err := client.Set(ctx, "key", "value", 0).Err()
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
}

func hashes() {
	fmt.Println("\nhashes started.======")
	client.HSet(ctx, "myHash", "field1", "value1", "field2", "value2")
	field1, err := client.HGet(ctx, "myHash", "field1").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("field1:", field1)

	myHash, err := client.HGetAll(ctx, "myHash").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("myHash:", myHash)

	_ = client.HDel(ctx, "myHash", "key2")
}

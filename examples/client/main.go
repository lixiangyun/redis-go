package main

import (
	"context"
	"fmt"
	"time"

	"redis-go"
)

func main() {

	fmt.Println(context.Background())

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	redis.DefaultClient = &redis.Client{Addr: "192.168.0.107:6379"}

	// Use the default client which is configured to connect to the redis server
	// running at localhost:6379.
	if err := redis.Exec(ctx, "SET", "hello", "world"); err != nil {
		fmt.Println(err)
	}

	// Passing request or response arguments is done by consuming the stream of
	// values.
	var args = redis.Query(ctx, "GET", "hello")
	var value string

	if args.Next(&value) {
		fmt.Println(value)
	}

	if err := args.Close(); err != nil {
		fmt.Println(err)
	}
}

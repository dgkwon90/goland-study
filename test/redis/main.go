package main

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

func main() {
	// New Client
	rc := redis.NewClient(&redis.Options{
		Addr:     "192.168.56.1:6379",
		Password: "test001", // password
		DB:       0,         // namespace
	})

	ctx := context.Background()
	result, err := rc.Ping(ctx).Result()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Connect: %v\n", result)

	err2 := rc.Set(ctx, "test", "1234", 0).Err()
	if err2 != nil {
		return
	}
	fmt.Printf("Set: %v -> %v\n", "test", "1234")

	val, err3 := rc.Get(ctx, "test").Result()
	if err3 != nil {
		return
	}
	fmt.Printf("Get: %v -> %v\n", "test", val)

}

package main

import (
	"fmt"
	"context"
	"github.com/go-redis/redis/v8"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	rdb := redis.NewFailoverClient(&redis.FailoverOptions{
		MasterName:    "mymaster",
		SentinelAddrs: []string{"localhost:26379"},
	})
	//x :=rdb.Ping(ctx)

	rdb.Do(ctx, "RPUSH", "HELLO", "WORLD")

	reply := rdb.Do(ctx, "BLPOP", "HELLO", 5)
	//fmt.Println(x)
	//x := reply.Val()

	x, _ := reply.Result()

	ss := x.([]interface{})[1]

	fmt.Println(ss)
}

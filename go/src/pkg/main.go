package main

import (
	"fmt"
	"pkg/components"
)

func main() {
	redis := components.NewRedisClient()
	key := "name"
	err := redis.Set(key, "jack", 0).Err()
	if err != nil {
		panic(err)
	}

	val, _ := redis.Get(key)
	fmt.Println(val)
}

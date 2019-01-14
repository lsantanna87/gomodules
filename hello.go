package main

import (
	"fmt"
	"log"

	"github.com/lsantanna87/gomodules/redis"
	"github.com/mediocregopher/radix/v3"
)

func main() {
	var err error
	redis.Client, err = radix.NewCluster([]string{"127.0.0.1:7000"})
	defer redis.Client.Close()

	if err != nil {
		log.Fatal("Error when trying to get redis connection")
	}
	redis.Set("foo", "lucas")
	fmt.Println(redis.Get("foo"))
}

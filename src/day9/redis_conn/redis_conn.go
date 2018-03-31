package main

import (
	"github.com/garyburd/redigo/redis"
	"fmt"
)

func main()  {

	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("error")
		return
	}
	defer conn.Close()

	conn.Do("set", "a", "c")


}

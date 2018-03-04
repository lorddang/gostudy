package main

import (
	"time"
	"fmt"
)

const (
	man = 1
	female = 2
)

func main()  {
	second := time.Now().Unix()
	if second % female == 0 {
		fmt.Println("female")
	}else {
		fmt.Println("man")
	}
}
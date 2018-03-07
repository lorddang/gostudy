package main

import (
	"fmt"
	"time"
)

func main()  {
	printN (100)
	time.Sleep(time.Second)

}

func printN (n int) {

	for i := 0; i <= n ; i++ {
		go fmt.Printf("%d + %d = %d \n" , i , n - i, n)
	}
}

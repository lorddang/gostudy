package main

import (
	"fmt"
)

func main()  {
	interChannel := make(chan int, 1000)
	exitChan := make(chan bool, 2)
	go generageInteger(interChannel, exitChan, 1000)
	go printInteger(interChannel, exitChan)
	<- exitChan
	<- exitChan
	fmt.Println("over")
}

func generageInteger(c chan int,ec chan bool, n int)  {
	for i := 0; i< n; i++ {
		c <- i
	}
	close(c)
	ec <- true
	fmt.Println("g over")
}

func printInteger(c chan int, ec chan bool)  {
	for i := range c{
		fmt.Println(i)
	}
	fmt.Println("go over")
	ec <- true
}
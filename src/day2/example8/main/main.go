package main

import(
	"fmt"
)

//写一个程序用来打印值类型和引用类型变量到终端，并观察输出结果

func main()  {
	var a = 10
	var channel chan int
	channel = make(chan int , 3)
	channel <- 3
	fmt.Printf("%p \n", channel)
	b := <- channel
	fmt.Printf("%d \n", a)
	fmt.Printf("%d \n", b)
}
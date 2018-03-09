package main

import(
	"math/rand"
	"fmt"
)

func main()  {
	var intChannel = make(chan int, 10)
	var floatChannel = make(chan float32, 10)
	fmt.Println("=======整数==========")
	for i := 0;i< 10; i++  {
		go genRandomInt(intChannel)
		x := <- intChannel
		fmt.Println(x)
	}
	fmt.Println("======小于10的整数=======")
	for i := 0;i< 10; i++  {
		go genRandomIntLess100(intChannel)
		x := <- intChannel
		fmt.Println(x)
	}
	fmt.Println("========浮点数===========")
	for i := 0; i < 10; i ++ {
		go genRandomFloat(floatChannel)
		f := <- floatChannel
		fmt.Println(f)
	}


}
// 生成随机数并放入管道
func genRandomInt(c chan int ) {
	c <- rand.Int()
}
// 生成浮点数并放入管道
func genRandomFloat(c chan float32) {
	c <- rand.Float32()
}

// 生成小于10的随机整数
func genRandomIntLess100(c chan int){
		c <- rand.Intn(10)
}

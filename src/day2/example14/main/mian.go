package main

import(
	"math/rand"
	"fmt"
	"time"
)

func main()  {
	var intChannel = make(chan int, 10)
	var floatChannel = make(chan float32, 10)
	for i := 0;i< 10; i++  {
		go genRandomInt(intChannel)
		x := <- intChannel
		fmt.Println(x)
	}
	fmt.Println("==============================")
	for i := 0; i < 10; i ++ {
		go genRandomFloat(floatChannel)
		f := <- floatChannel
		fmt.Println(f)
	}

	time.Sleep(time.Second * 1)

}

func genRandomInt(c chan int ) {
	c <- rand.Int()
}

func genRandomFloat(c chan float32) {
	c <- rand.Float32()
}

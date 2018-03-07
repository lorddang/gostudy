package main

import(
	"fmt"
	"math"
)


func main()  {
	var a  int
	a = 10
	b := math.Sin(float64(a))
	fmt.Printf("%x \n", a)
	fmt.Printf("%f \n", b)


}
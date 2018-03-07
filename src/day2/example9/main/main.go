package main


import(
	"fmt"
)

func main()  {
	a := 1
	b := 2

	a, b = exchangeValue(a, b)
	fmt.Printf("a=%d, b=%d", a, b)
}
// 交换两个变量的值
func exchangeValue(a int, b int) (int, int)  {

	var c int
	c = a
	a = b
	b = c
	return a, b

}
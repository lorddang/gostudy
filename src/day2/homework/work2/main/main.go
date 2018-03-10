/*
打印出100-999中所有的“水仙花数”，所谓“水仙花数”是指一个三位数，其各位数字
立方和等于该数本身。例如：153 是一个“水仙花数”，因为 153=1 的三次
方＋5 的三次方＋3 的三次方。
*/
package main

import (
	"fmt"
	"math"
)

func main() {
	for i := 100; i < 1000; i++ {
		x := i / 100
		y := (i - 100*x) / 10
		n := i - 100*x - 10*y
		if int(math.Pow(float64(x), 3)+math.Pow(float64(y), 3)+math.Pow(float64(n), 3)) == i {
			fmt.Println(i)
		}
	}

}

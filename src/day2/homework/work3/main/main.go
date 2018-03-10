// 对于一个数n，求n的阶乘之和，即： 1！ + 2！ + 3！+…n!
package main

import "fmt"

func main() {
	var n int = 10
	var result float64
	for i := 1; i <= n; i++ {
		result = result + n_n(i)
	}
	fmt.Printf("%d 的阶乘之和为 %f", n, result)
}

func n_n(n int) float64 {
	var result float64 = 1
	for i := 1; i < n; i++ {
		result = result * float64(i)
	}
	return result
}

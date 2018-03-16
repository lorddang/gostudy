// 使用递归调用的方式生成斐波那契数列
package main

import "fmt"

func main()  {

	for i:= 1; i <= 10; i++ {
		fmt.Println(fab(i))
	}

}

func fab(n int) int64 {
	if n == 1 || n == 2 {
		return  1
	}

	return fab(n - 1) + fab(n - 2)

}
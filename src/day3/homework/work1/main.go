// 编写程序，在终端输出九九乘法表
package main

import "fmt"

func main()  {

	for i := 1; i <= 9; i ++ {
		for j := 1; j<=i; j++ {
			fmt.Printf("%d x %d = %d   ", i, j, i * j)
		}
		fmt.Println()
	}
}

// 使用切片的方式生成斐波纳挈数列
package main

import (
	"fmt"
)

func main() {
	for _, v := range fab(50) {
		fmt.Println(v)
	}
}

func fab(n int) []int64 {

	fab := make([]int64, 2)
	fab[1] = 1
	fab[0] = 1
	for i := 2; i < n; i++ {
		fab = append(fab, fab[i-1]+fab[i-2])
	}
	return fab
}

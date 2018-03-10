/*
判断101-200 之间有多少个素数，并输出所有的素数
*/
package main

import (
	"fmt"
	"math"
)

func main() {
	for i := 101; i <= 200; i++ {
		isSuShu := true
		for j := 2; j < int(math.Sqrt(float64(i))); j++ {
			if i%j == 0 {
				isSuShu = false
				break

			}
		}
		if isSuShu == true {
			fmt.Println(i)
		}
	}
}
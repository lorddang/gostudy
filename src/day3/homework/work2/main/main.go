// 一个数如果恰好等于它的因子之和，这个数就称为“完数”。例如6=1＋2＋3. 编程找出1000以内的所有完数。

package main

import "fmt"

func main()  {
	for i:=1; i<=1000;i++ {
		if i == sumOfFactor(i) {
			fmt.Println(i)
		}
	}


}

func sumOfFactor(n int) int{
	sum := 1
	for i := 2; i <= n / 2 ; i++ {
		if n % i == 0 {
			sum += i
		}
	}
	return sum
}
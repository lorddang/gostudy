// 输入一行字符，分别统计出其中英文字母、空格、数字和其它字符的个数。
package main

import (
	"fmt"
)

func main()  {
	var str string
	fmt.Println("请输入字符串")
	fmt.Scanf("%s\n", &str)
	numOfSpace, numOfEnglishChar, numOfNumber, numOfOthers := countByCategory(str)
	fmt.Printf("空格的个数为：%d \n", numOfSpace)
	fmt.Printf("英文字母的个数为：%d \n", numOfEnglishChar)
	fmt.Printf("数字的个数为：%d \n", numOfNumber)
	fmt.Printf("其他字符的个数为：%d \n", numOfOthers)
}

func countByCategory(s string) (int, int, int, int)  {
	str := []rune(s)
	var (
		numOfSpace int
		numOfEnglishChar int
		numOfNumber int
		numOfOthers int
	)
	for i := 0; i < len(str); i ++ {
		if str[i] == ' ' {
			numOfSpace ++
		} else if (str[i] >= 'a' && str[i] <= 'z') || (str[i] >= 'A' && str[i] <= 'Z') {
			numOfEnglishChar ++
		} else if str[i] > 0 && str[i] < 9 {
			numOfNumber ++
		} else {
			numOfOthers ++
		}
	}


	return numOfSpace, numOfEnglishChar, numOfNumber, numOfOthers
}

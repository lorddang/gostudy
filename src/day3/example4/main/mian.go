// 写一个函数分别演示Replace、Count、Repeat、ToLower、ToUpper 的用法
package main

import (
	"fmt"
	"strings"
)

func main()  {
	 replace()
	 count()
	 repeat()
	 upperAndLower()

}

func replace()  {
	var str, replaceStr, newStr string
	fmt.Println("Please input a string")
	fmt.Scanf("%s\n", &str)
	fmt.Println("Please input which want to replace string and new string")
	fmt.Scanf("%s%s\n", &replaceStr, &newStr)
	result := strings.Replace(str, replaceStr, newStr, -1)
	fmt.Printf("string is %s, replace string is %s , new string is %s , result is %s ", str, replaceStr, newStr, result)
}

func count()  {
	var str, countStr string
	fmt.Println("Please input string")
	fmt.Scanf("%s\n", &str)
	fmt.Println("Please input which want to count str")
	fmt.Scanf("%s\n", &countStr)
	fmt.Printf("%s appears %d times in %s", countStr, strings.Count(str, countStr), str)
}

func repeat()  {
	var str string
	var count int
	fmt.Println("Please input string and repeat times")
	fmt.Scanf("%s%d\n", &str, &count)
	fmt.Println(strings.Repeat(str, count))
}

func upperAndLower()  {
	var str string
	fmt.Println("Please input string")
	fmt.Scanf("%s\n", &str)
	fmt.Printf("Upper is %s, Lower is %s", strings.ToUpper(str), strings.ToLower(str))
}

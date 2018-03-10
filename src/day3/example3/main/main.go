// 编写一个函数返回一个字符串在另一个字符串的首次和最后次出现的位置
package main

import (
	"fmt"
	"strings"
)

func strIndex(str string, substr string) (f, l int)  {

	f = strings.Index(str, substr)
	l = strings.LastIndex(str, substr)
	return
}

func main()  {
	var str, sub string
	fmt.Scanf("%s%s\n", &str, &sub)

	first, last := strIndex(str, sub)
	fmt.Printf("first Position is %d , last Position is %d", first, last)

}

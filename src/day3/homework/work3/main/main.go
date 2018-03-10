// 输入一个字符串，判断其是否为回文。回文字符串是指从左到右读和从右到
// 左读完全相同的字符串
package main

import "fmt"

func main()  {
	var str string
	fmt.Printf("请输入一个字符串：")
	fmt.Scanf("%s\n", &str)
	if isMirrorStr(str) {
		fmt.Printf("%s 是一个回文字符串", str)
	} else {
		fmt.Printf("%s 不是一个回文字符串", str)
	}
}

func isMirrorStr(str string) bool  {
	l := len(str)
	for i:=0;i<l/2;i++ {
		if str[i] != str[l - i -1] {
			return false
		}
	}
	return true
}
// 判断一个路径是否以/结尾,如果不是加上/。
package main

import (
	"fmt"
	"strings"
)

func main()  {
	var str string
	fmt.Scanf("%s", &str)

	if !strings.HasSuffix(str, "/"){
		str = str + "/"
	}
	fmt.Println(str)
}

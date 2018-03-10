// 判断一个url是否以http:// 开头，如果不是,则加上http://.
package main

import (
	"fmt"
	"strings"
)

func main()  {
	var str string
	fmt.Scanf("%s", &str)

	if !strings.HasPrefix(str, "http://") {
		str = "http://" + str
	}
	fmt.Println(str)
}

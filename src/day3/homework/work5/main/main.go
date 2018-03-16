// 模拟两个大数加
package main

import (
	"bufio"
	"os"
	"strings"
	"fmt"
)

func main()  {
	reader := bufio.NewReader(os.Stdin)
	line, _, err := reader.ReadLine()
	str := string(line)
	if err != nil {
		fmt.Printf("enconter an error !!!")
		return
	}
	result := strings.Split(str, "+")
	str1 := strings.TrimSpace(result[0])
	str2 := strings.TrimSpace(result[1])
	fmt.Println(plus(str1, str2))

}

func plus(str1, str2 string) (result string)  {

	s1 := str1
	s2 := str2
	if len(str1) < len(str2) {
		s1 = str2
		s2 = str1
	}
	index1 := len(s1) - 1
	index2 := len(s2) - 1
	var left int
	for index1 >=0 && index2 >= 0 {

		c1 := s1[index1] - '0'
		c2 := s2[index2] - '0'
		sum := int(c1) + int(c2) + left
		left = int(sum / 10)
		result = fmt.Sprintf("%c%s", sum % 10 + '0', result)
		index1 --
		index2 --
	}
	for index1 >= 0 {
		c1 := s1[index1] - '0'
		sum := int(c1) + left
		left = int(sum / 10)
		fmt.Println(left)
		result = fmt.Sprintf("%c%s", sum % 10 + '0', result )
		index1 --

	}

	if left == 1 {
		result = fmt.Sprintf("1%s", result)
	}
	return
}
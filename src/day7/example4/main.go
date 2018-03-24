package main

import (
	"fmt"
	"bufio"
	"os"
)

func main()  {
	fmt.Print("请输入: ")
	reader := bufio.NewReader(os.Stdin)
	lineString, _ := reader.ReadString('\n')
	numOfInt, numOfChar, numOfSpace, numOfOther := processString(lineString)
	fmt.Printf("数字: %d \n", numOfInt)
	fmt.Printf("字符: %d \n", numOfChar)
	fmt.Printf("空格: %d \n", numOfSpace)
	fmt.Printf("其他: %d \n", numOfOther)
}

func processString(str string) (numOfInt, numOfChar, numOfSpace, numOfOther int)  {
	for _, c := range []rune(str){
		switch {
		case c >= '0' && c <= '9':
			numOfInt ++

		case c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'z':
			numOfChar ++
		case c == ' ':
			numOfSpace ++
		default:
			numOfOther ++
		}
	}
	return
}
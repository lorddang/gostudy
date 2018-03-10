package main

import (
	"fmt"
	"strings"
)

func main()  {
	trimSpace()
}


func trimSpace(){
	str := ` jidfji jiojodif joidjof`
	fmt.Printf("Trimed string is %s", strings.TrimSpace(str))
}
package main

import (
	"fmt"
)
type myprint interface {
	print()
}

type Mystruce struct {
	name string
}

func (m *Mystruce)print()  {
	fmt.Println(m.name)
}
func main()  {
	var a interface{}
	var b int
	b = 1
	a = b
	print(a, b, "String")
}

func print(i ...interface{})  {
	for _, value := range i{
		fmt.Printf("%v\n", value)
		switch value.(type) {
		case int:
			fmt.Println("int")
		case string:
			fmt.Println("string")
		}
	}
	var a myprint
	a = &Mystruce{}
	if _, ok := a.(myprint)  ; ok{
		fmt.Println(ok)
	}
}
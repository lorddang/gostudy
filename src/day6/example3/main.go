package main

import (
	"reflect"
	"fmt"
)

func main()  {
	reflectTest()
	return
}

func reflectTest()  {
	var a interface{}
	a = 100
	v := reflect.ValueOf(a)
	fmt.Println(v.Type(), v.Kind())

}

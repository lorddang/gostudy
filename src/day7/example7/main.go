package main

import (
	"flag"
	"fmt"
)

func main()  {
	process()
}

func process()  {
	var b string
	flag.StringVar(&b, "str", "test", "this is string test")
	flag.Parse()
	fmt.Println(b)
}

package main

import (
	"fmt"
	"os"
)


func main() {
	goRoot := os.Getenv("GOPATH")
	path := os.Getenv("PATH")

	fmt.Println(goRoot)
	fmt.Println(path)
}


package main

import (
	"fmt"
	"os"
	"strings"
)


func main() {
	goRoot := os.Getenv("GOPATH")
	path := os.Getenv("PATH")

	fmt.Println(goRoot)
	fmt.Println(path)
	fmt.Print(strings.Split(path, ":")[2])
}


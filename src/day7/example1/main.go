package main

import (
	"os"
	"fmt"
	"io"
)

func main()  {
	fileReader()
}

func fileReader()  {
	file, err := os.Open("/Users/abner/Documents/tmp/api.sql")
	defer file.Close()
	if err != nil {
		fmt.Println("error")
	}
	outfile, outErr := os.OpenFile("/Users/abner/Documents/tmp/api.sql.bak", os.O_WRONLY|os.O_CREATE, 0666)
	if outErr != nil{
		fmt.Println("writer error")
	}
	defer outfile.Close()
	io.Copy(outfile, file)
}

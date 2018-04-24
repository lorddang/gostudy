package main

import (
	"net/http"
	"fmt"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/table", table)
	err := http.ListenAndServe("0.0.0.0:8000", nil)
	if err != nil{
		fmt.Println(err)
	}
}

func index(respw http.ResponseWriter, req *http.Request)  {
	fmt.Fprintf(respw, "hello world" )
}

func table(respw http.ResponseWriter, req *http.Request)  {

	fmt.Fprintf(respw, "helloworld %d", 1)
}

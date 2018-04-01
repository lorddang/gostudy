package main

import (
	"net/http"
	"fmt"
)

func main() {
	http.HandleFunc("/", index)
	err := http.ListenAndServe("0.0.0.0:80", nil)
	if err != nil{
		fmt.Println("error")
	}
}

func index(respw http.ResponseWriter, req *http.Request)  {
	fmt.Println(req.Header["User-Agent"])
	fmt.Fprintf(respw, "<h1>It's Works</h1>")
}

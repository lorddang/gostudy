package main

import (
	"context"
	"time"
	"net/http"
	"fmt"
	"io/ioutil"
)

type Result struct {
	r *http.Response
	err error
}


func process() {
	ctx , cancle := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancle()
	tr := &http.Transport{}
	clent := http.Client{Transport:tr,}
	c := make(chan Result, 1)
	req ,err := http.NewRequest("GET", "http://www.google.com", nil)
	if err != nil{
		fmt.Println(err)
	}
	go func() {
		resp, err := clent.Do(req)
		pack := Result{r: resp, err: err,}
		c <- pack
	}()
	select {
	case <- ctx.Done():
		tr.CancelRequest(req)
		res := <-c
		fmt.Print("Time Out")
		fmt.Println(res.r, res.err)
	case res := <- c:
		defer res.r.Body.Close()
		out, _ := ioutil.ReadAll(res.r.Body)
		fmt.Printf("%s", out)
	}

}


func main()  {
	process()
}

package main

import (
	"mysql"
	"fmt"
)
func main()  {

	driver := mysql.MySQLDriver{}
	client, err := driver.Open("root:@tcp(127.0.0.1)/")
	if err != nil {
		fmt.Println(err)
		return
	}
	// defer client.Close()
	if client.RegisterSlave() == nil {
		fmt.Println("注册成功")
	} else {
		fmt.Println("注册失败")
		return
	}
	//return
	data := client.StartSlaveIOThread()
	fmt.Println(data)


}
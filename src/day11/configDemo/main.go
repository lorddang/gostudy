package main

import (
	"github.com/astaxie/beego/config"
	"fmt"
)

func main() {
	conf, err := config.NewConfig("ini", "src/day11/configDemo/config.ini")
	if err!= nil {
		fmt.Printf("new config error %v\n", err)
		return
	}
	logLevel := conf.String("log::log_level")
	logPath := conf.String("log::log_path")

	fmt.Println("logLevel:", logLevel)
	fmt.Println("logPath:", logPath)

}

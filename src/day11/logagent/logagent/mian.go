package main

import (
	"day11/logagent/conf"
	"fmt"
	"github.com/astaxie/beego/logs"
	"day11/logagent/tailf"
	"day11/logagent/kafka"
)
func main()  {
	err := conf.InitConfig("ini", "./config.ini")
	if err != nil {
		fmt.Printf("init confi error %v \n", err)
		return
	}
	err = initLogs()
	if err != nil {
		fmt.Printf("init logs failed %v", err)
		return
	}else {
		logs.Info("init logs success !!!")
	}

	err = tailf.InitTail(conf.CollConfs)

	kafkaClient, err := kafka.InitKafka(conf.KafkaConf)
	if err != nil {
		logs.Error("init kafka failed ", err)
		return
	}else {
		logs.Info("init kafka success !!!")
	}
	err = server(kafkaClient)
	if err != nil {
		logs.Error("send msg error ", err)
		return
	}else {
		logs.Info("logagent exit")
	}


}
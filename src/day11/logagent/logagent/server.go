package main

import (
	"github.com/Shopify/sarama"
	"day11/logagent/tailf"
	"github.com/astaxie/beego/logs"
)

func server(kc sarama.SyncProducer) error  {
	for {
		lmsg := tailf.ReadFromChan()
		msg := &sarama.ProducerMessage{}
		msg.Topic = lmsg.Topic
		msg.Value = sarama.StringEncoder(lmsg.Msg)
		pid, offset, err := kc.SendMessage(msg)
		if err != nil {
			logs.Error("send message failed,", err)
			return err
		}

		logs.Info("send log success pid:%v offset:%v", pid, offset)
	}
	return nil
}
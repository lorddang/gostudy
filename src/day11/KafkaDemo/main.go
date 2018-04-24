package main

import (
	"fmt"
	"github.com/Shopify/sarama"
)

func main() {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true
	client, err := sarama.NewConsumer([]string{"zk1:9092"}, config)
	if err != nil {
		fmt.Println("producer close, err:", err)
		return
	}
	defer client.Close()
	pc , err :=client.ConsumePartition("niginx_log", 3, 0)
		for msg := range pc.Messages(){
			fmt.Println(msg.Value)
	}


	//msg := &sarama.ProducerMessage{}
	//msg.Topic = "nginx_log"
	//msg.Value = sarama.StringEncoder("hello world")
	//client, err := sarama.NewSyncProducer([]string{"zk1:9092"}, config)
	//if err != nil {
	//	fmt.Println("producer close, err:", err)
	//	return
	//}
	//
	//defer client.Close()
	//
	//pid, offset, err := client.SendMessage(msg)
	//if err != nil {
	//	fmt.Println("send message failed,", err)
	//	return
	//}
	//
	//fmt.Printf("pid:%v offset:%v\n", pid, offset)
}

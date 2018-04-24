package main

import (
	"github.com/Shopify/sarama"
	"fmt"
	"sync"
)

var (
	wg sync.WaitGroup
)

func main() {
	config := sarama.NewConfig()
	client, err := sarama.NewConsumer([]string{"zk1:9092"}, config)
	if err != nil {
		fmt.Println("producer close, err:", err)
		return
	}
	defer client.Close()
	pl, err := client.Partitions("nginx_log")
	for _, p := range pl {
		pc , err := client.ConsumePartition("nginx_log", p, 121053)
		if err != nil {
			fmt.Println("close, err:", err)
			return
		}
		wg.Add(1)
		go func(pc sarama.PartitionConsumer) {
			for msg := range pc.Messages(){
				fmt.Println(string(msg.Value))
			}
			wg.Done()
		}(pc)
	}
	wg.Wait()

}

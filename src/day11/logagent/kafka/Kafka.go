package kafka

import (
	"github.com/Shopify/sarama"
	"fmt"
	"day11/logagent/conf"
)

func InitKafka(conf *conf.KfkConf) (kafkaClient sarama.SyncProducer, err error) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true


	kafkaClient, err = sarama.NewSyncProducer([]string{conf.ServerAddr}, config)
	if err != nil {
		fmt.Println("producer close, err:", err)
		return
	}
	return

}



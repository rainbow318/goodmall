package mq

import (
	"github.com/IBM/sarama"
	"github.com/suutest/app/checkout/conf"
)

var SaramaClient *sarama.Client

func InitSaramaClient() {
	config := sarama.NewConfig()
	config.Producer.Partitioner = sarama.NewRandomPartitioner // // 新选出一个partition
	config.Producer.Return.Successes = true                   // 成功交付的消息将在success channel返回

	// 连接kafka
	SaramaClient, err := sarama.NewSyncProducer(conf.GetConf().Kafka.Address, config)
	if err != nil {
		panic(err)
	}
	defer SaramaClient.Close()
}

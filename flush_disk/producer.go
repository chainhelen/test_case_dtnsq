package main

import (
	"fmt"
	dtnsq "github.com/chainhelen/go-dtnsq"
)

func NewMQProducer() *dtnsq.Producer {
	cfg := dtnsq.NewConfig()
	cfg.Set("agreement_version", "DT")
	MQProducer, err := dtnsq.NewProducer("127.0.0.1:4150", cfg)
	if err != nil {
		panic(fmt.Sprintf("new mq producer fail, err: %v", err))
	}
	return MQProducer
}
func main() {
	producer := NewMQProducer()

	if err := producer.Publish("flush_topic", []byte("hello dtnsq one")); err != nil {
		fmt.Printf("pre:%s\n", err.Error())
	}
}

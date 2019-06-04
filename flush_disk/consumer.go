package main

import (
	"fmt"
	dtnsq "github.com/chainhelen/go-dtnsq"
)

func main() {
	cfg := dtnsq.NewConfig()
	cfg.Set("agreement_version", "DT")
	MQConsumer, err := dtnsq.NewConsumer("flush_topic", "ch1", cfg)
	if err != nil {
		panic(fmt.Sprintf("new mq consumer fail, err: %v", err))
	}

	// 设置消息处理函数
	MQConsumer.AddHandler(dtnsq.HandlerFunc(func(message *dtnsq.Message) error {
		fmt.Printf("id:%d, body:%s\n", message.ID, message.Body)
		return nil
	}))

	err = MQConsumer.ConnectToNSQD("127.0.0.1:4150")
	if err != nil {
		panic(fmt.Sprintf("conect mq consumer fail, err: %v", err))
	}
	<-MQConsumer.StopChan
}

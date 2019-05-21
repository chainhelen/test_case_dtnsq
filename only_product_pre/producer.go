package main

import (
	"fmt"
	dtnsq "github.com/chainhelen/go-dtnsq"
	"time"
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
	producer.SetCheckBackHandler(func(msg *dtnsq.Message) error {
		fmt.Printf("Producer receive checkback message %s\n\n", msg)
		time.Sleep(time.Duration(3) * time.Second)
		return nil
	})

	// pre、wait 3s、cmt
	func() {
		if preres, err := producer.PublishDtPre("only_product_topic", []byte("hello dtnsq one")); err != nil {
			fmt.Printf("pre err:%s\n", err.Error())
		} else {
			time.Sleep(time.Duration(3) * time.Second)
			fmt.Printf("pre res:%s\n", preres)
		}
	}()

	// pre、wait 3s、cnl
	func() {
		if preres, err := producer.PublishDtPre("only_product_topic", []byte("hello dtnsq two")); err != nil {
			fmt.Printf("pre err:%s\n", err.Error())
		} else {
			time.Sleep(time.Duration(3) * time.Second)
			fmt.Printf("pre res:%s\n", preres)
		}
	}()

	time.Sleep(time.Hour * time.Duration(3))
}

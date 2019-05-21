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

	// pre、wait 3s、cmt
	func() {
		if preres, err := producer.PublishDtPre("normal_topic", []byte("hello dtnsq one")); err != nil {
			fmt.Printf("pre:%s\n", err.Error())
		} else {
			time.Sleep(time.Duration(3) * time.Second)
			fmt.Printf("pre:%s\n", preres)
			if cmtres, err := producer.PublishDtCmt("normal_topic", preres); err != nil {
				fmt.Printf("cmt:%s\n", err.Error())
			} else {
				fmt.Printf("cmt:%s\n", cmtres)
			}
		}
	}()

	// pre、wait 3s、cnl
	func() {
		if preres, err := producer.PublishDtPre("normal_topic", []byte("hello dtnsq tow")); err != nil {
			fmt.Printf("pre:%s\n", err.Error())
		} else {
			time.Sleep(time.Duration(3) * time.Second)
			fmt.Printf("pre:%s\n", preres)
			if cmtres, err := producer.PublishDtCnl("normal_topic", preres); err != nil {
				fmt.Printf("cnl:%s\n", err.Error())
			} else {
				fmt.Printf("cnl:%s\n", cmtres)
			}
		}
	}()
}

package main

import (
	"fmt"
	dtnsq "github.com/chainhelen/go-dtnsq"
	"time"
)

func NewMQProducer() *dtnsq.Producer {
	cfg := dtnsq.NewConfig()
	//	cfg.Set("agreement_version", "DT")
	MQProducer, err := dtnsq.NewProducer("127.0.0.1:4150", cfg)
	if err != nil {
		panic(fmt.Sprintf("new mq producer fail, err: %v", err))
	}
	return MQProducer
}
func main() {
	producer := NewMQProducer()

	// pre、wait 3s、cmt
	if err := producer.Publish("only_product_topic", []byte("hello dtnsq one")); err != nil {
		fmt.Errorf("%s\n", err.Error())
	}
	time.Sleep(time.Duration(3) * time.Hour)

}

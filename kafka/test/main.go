package main

import (
	"github.com/Shopify/sarama"
	"log"
)

func main() {
	config := sarama.NewConfig()
	config.Producer.Return.Errors = true
	w, _ := sarama.NewAsyncProducer([]string{"127.0.0.1:9092"}, config)

	log.Println("hongbao kfk broker writer runHongbaoBroker")
	defer func() {
		log.Println("hongbao kfk broker writer stop")
		w.Close()
	}()

	kmsg := &sarama.ProducerMessage{
		Topic: "web_log",
		Value: sarama.ByteEncoder("bb"),
	}
	w.Input() <- kmsg
}

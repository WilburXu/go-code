package main

import (
	"github.com/Shopify/sarama"
	"log"
	"time"
)

func main() {
	config := sarama.NewConfig()
	config.Producer.Return.Errors = true
	config.Producer.Return.Successes = true
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Version = sarama.V1_0_0_0

	producer, _ := sarama.NewAsyncProducer([]string{"127.0.0.1:9092"}, config)

	log.Println("hongbao kfk broker writer runHongbaoBroker")
	defer func() {
		log.Println("hongbao kfk broker writer stop")
		producer.Close()
	}()


	for {
		log.Println("WilburXu start send msg")

		kmsg := &sarama.ProducerMessage{
			Topic: "web_log",
			Value: sarama.ByteEncoder("WilburXu send msg"),
		}
		producer.Input() <- kmsg
		select {
		case suc := <-producer.Successes():
			log.Printf("> message: \"%s\" sent to partition  %d at offset %d\n", suc.Value,  suc.Partition, suc.Offset)
		case fail := <-producer.Errors():
			log.Printf("err: %s\n", fail.Err.Error())
		default:
			log.Println("???? why?")
		}

		time.Sleep(time.Second * 1)
	}
}

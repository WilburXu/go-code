package main

import (
	"github.com/Shopify/sarama"
	"log"
)

func main() {
	log.Println("consumer_test")

	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	config.Version = sarama.V0_11_0_2

	// consumer
	consumer, err := sarama.NewConsumer([]string{"localhost:9092"}, config)
	if err != nil {
		log.Println("consumer_test create consumer error %s\n", err.Error())
		return
	}

	defer consumer.Close()

	partition_consumer, err := consumer.ConsumePartition("web_log", 0, sarama.OffsetOldest)
	if err != nil {
		log.Println("try create partition_consumer error %s\n", err.Error())
		return
	}
	defer partition_consumer.Close()

	for {
		select {
		case msg := <-partition_consumer.Messages():
			log.Println("msg offset: %d, partition: %d, timestamp: %s, value: %s\n",
				msg.Offset, msg.Partition, msg.Timestamp.String(), string(msg.Value))
		case err := <-partition_consumer.Errors():
			log.Println("err :%s\n", err.Error())
		}
	}

}

//package main
//
//import (
//	"context"
//	"encoding/json"
//	"log"
//	"strings"
//	"time"
//
//	"github.com/Shopify/sarama"
//)
//
//// kafka consumer
//
//type LiveUpSt struct {
//	Data interface{} `json:"data"`
//}
//
//type Poster struct {
//}
//
//var DefaultPoster *Poster
//
//func main() {
//	DefaultPoster = &Poster{}
//	config := sarama.NewConfig()
//	config.Consumer.Return.Errors = true
//	config.Version = sarama.V1_0_0_0
//	config.Consumer.Offsets.Initial = sarama.OffsetNewest
//
//	r, err := sarama.NewConsumerGroup(strings.Split("127.0.0.1:9092", ","), "test", config)
//	if err != nil {
//		panic(err.Error())
//	}
//	for {
//		err := r.Consume(context.Background(), strings.Split("web_log", ","), DefaultPoster)
//		if err != nil {
//			time.Sleep(100 * time.Millisecond)
//			continue
//		}
//	}
//}
//
//func (Poster) Setup(sarama.ConsumerGroupSession) error   { return nil }
//func (Poster) Cleanup(sarama.ConsumerGroupSession) error { return nil }
//func (m Poster) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
//	for msg := range claim.Messages() {
//		sess.MarkMessage(msg, "")
//
//		log.Println("kfk broker recv msg: %s", string(msg.Value))
//
//		st := LiveUpSt{}
//		err := json.Unmarshal(msg.Value, &st)
//		if err != nil {
//			log.Println("failed to parsre mqmsg: %s - %s", string(msg.Value), err.Error())
//			continue
//		}
//
//		log.Printf("st ret %v \n", st)
//	}
//	return nil
//}

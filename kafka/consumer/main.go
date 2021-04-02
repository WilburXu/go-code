//package main
//
//import (
//	"github.com/Shopify/sarama"
//	"log"
//	"strings"
//	"time"
//)
//
//func main() {
//	log.Println("consumer_test")
//
//	config := sarama.NewConfig()
//	config.Consumer.Return.Errors = true
//	config.Version = sarama.V0_11_0_2
//
//	// consumer
//	consumer, err := sarama.NewConsumer([]string{"localhost:9092"}, config)
//	if err != nil {
//		log.Println("consumer_test create consumer error %s\n", err.Error())
//		return
//	}
//
//	defer consumer.Close()
//
//	partition_consumer, err := consumer.ConsumePartition("gift_send", 0, sarama.OffsetOldest)
//	if err != nil {
//		log.Println("try create partition_consumer error %s\n", err.Error())
//		return
//	}
//	defer partition_consumer.Close(
//
//	for {
//		select {
//		case msg := <-partition_consumer.Messages():
//			log.Println("msg offset: %d, partition: %d, timestamp: %s, value: %s\n",
//				msg.Offset, msg.Partition, msg.Timestamp.String(), string(msg.Value))
//		case err := <-partition_consumer.Errors():
//			log.Println("err :%s\n", err.Error())
//		}
//	}
//
//}

package main

import (
	"context"
	"encoding/json"
	"github.com/Shopify/sarama"
	"log"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"time"
)

/**
{
"id": 28537,
"mid": 10,
"order_type": 2,
"status": 1,
"channel": 2,
"ct": "2020-12-27T14:44:19.110290786+08:00",
"ut": "2020-12-27T14:44:19.110296122+08:00",
"to_mid": 10052,
"sid": 5001665,
"gift_id": "1e6f5043",
"count": 1,
"price": 49,
"is_lucky": 0,
"lucky_prize": 0,
"ext": null
}
 */

type KafkaGiftOrderSt struct {
	ID        int64     `json:"id"`
	Mid       int64     `json:"mid"`
	OrderType int8      `json:"order_type"`
	Channel   int16     `json:"channel"`
	Ct        time.Time `json:"ct"`
	Ut        time.Time `json:"ut"`
	ToMid     int64     `json:"to_mid"`
	Sid       int64     `json:"sid"`
	GiftID    string    `json:"gift_id"`
	Count     int64     `json:"count"`
	Price     int64     `json:"price"`
}

// Consumer represents a Sarama consumer group consumer
type Consumer struct {
	ready chan bool
}

func (c *Consumer) Setup(session sarama.ConsumerGroupSession) error {
	//panic("implement me")
	return nil
}

func (c *Consumer) Cleanup(session sarama.ConsumerGroupSession) error {
	//panic("implement me")
	return nil
}

func (c *Consumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		log.Printf("Message claimed: value = %s, timestamp = %v, topic = %s", string(message.Value), message.Timestamp, message.Topic)
		session.MarkMessage(message, "")

		st := KafkaGiftOrderSt{}
		_ = json.Unmarshal(message.Value, &st)
		log.Printf("st val %+v", st)
	}

	return nil
}

func main() {
	consumer := Consumer{
		ready: make(chan bool),
	}

	brokerList := []string{"127.0.0.1:9092"}

	config := sarama.NewConfig()
	config.Version = sarama.V1_0_0_0
	config.Consumer.Offsets.Initial = sarama.OffsetOldest

	group := "web_log_group_v3"

	ctx, cancel := context.WithCancel(context.Background())
	client, err := sarama.NewConsumerGroup(brokerList, group, config)
	if err != nil {
		log.Printf("Error creating consumer group client: %v \n", err)
	}

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			// `Consume` should be called inside an infinite loop, when a
			// server-side rebalance happens, the consumer session will need to be
			// recreated to get the new claims
			if err := client.Consume(ctx, strings.Split("account_gift_order_info,web_log", ","), &consumer); err != nil {
				log.Panicf("Error from consumer: %v", err)
			}
			// check if context was cancelled, signaling that the consumer should stop
			if ctx.Err() != nil {
				return
			}
			consumer.ready = make(chan bool)
		}
	}()

	<-consumer.ready // Await till the consumer has been set up
	log.Println("Sarama consumer up and running!...")

	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-ctx.Done():
		log.Println("terminating: context cancelled")
	case <-sigterm:
		log.Println("terminating: via signal")
	}
	cancel()
	wg.Wait()
	if err = client.Close(); err != nil {
		log.Panicf("Error closing client: %v", err)
	}

}

//-------------------
// ------------------
// 单一消费
//package main
//
//import (
//	"github.com/Shopify/sarama"
//	"log"
//	"os"
//	"os/signal"
//	"sync"
//	"syscall"
//)
//
//// kafka consumer
//
////var DefaultPoster *Poster
//
//var topic = "web_log"
//
//func main() {
//	config := sarama.NewConfig()
//	config.Consumer.Return.Errors = true
//	config.Consumer.Offsets.Initial = sarama.OffsetNewest
//
//	brokerList := []string{"127.0.0.1:9092"}
//	c, err := sarama.NewConsumer(brokerList, config)
//	//c, err := sarama.NewConsumerGroup(brokerList, "test", config)
//
//	partitionList, err := getPartitions(c, topic)
//	if err != nil {
//
//	}
//
//	log.Printf("partition %+v \n", partitionList)
//
//	var (
//		messages = make(chan *sarama.ConsumerMessage, 1024)
//		closing  = make(chan struct{})
//		wg       sync.WaitGroup
//	)
//
//	go func() {
//		signals := make(chan os.Signal, 1)
//		signal.Notify(signals, syscall.SIGTERM, os.Interrupt)
//		<-signals
//		log.Println("Initiating shutdown of consumer...")
//		close(closing)
//	}()
//
//
//	for _, partition := range partitionList {
//		pc, err := c.ConsumePartition(topic, partition, -1)
//		if err != nil {
//			log.Printf("Failed to start consumer for partition %d: %s \n", partition, err)
//		}
//
//		go func(pc sarama.PartitionConsumer) {
//			<-closing
//			pc.AsyncClose()
//		}(pc)
//
//		wg.Add(1)
//		go func(pc sarama.PartitionConsumer) {
//			defer wg.Done()
//			for message := range pc.Messages() {
//				messages <- message
//			}
//		}(pc)
//	}
//
//	go func() {
//		for msg := range messages {
//			log.Printf("Partition:\t%d\n", msg.Partition)
//			log.Printf("Offset:\t%d\n", msg.Offset)
//			log.Printf("Key:\t%s\n", string(msg.Key))
//			log.Printf("Value:\t%s\n", string(msg.Value))
//			log.Println()
//		}
//	}()
//
//	wg.Wait()
//	log.Println("Done consuming topic", topic)
//	close(messages)
//
//	if err := c.Close(); err != nil {
//		log.Println("Failed to close consumer: ", err)
//	}
//
//	//r, err := sarama.NewConsumerGroup(strings.Split("127.0.0.1:9092", ","), "test", config)
//	//if err != nil {
//	//	panic(err.Error())
//	//}
//	//for {
//	//	err := r.Consume(context.Background(), strings.Split("web_log", ","), DefaultPoster)
//	//	if err != nil {
//	//		time.Sleep(100 * time.Millisecond)
//	//		continue
//	//	}
//	//}
//}
//
//
//func getPartitions(c sarama.Consumer, topic string) ([]int32, error) {
//	return c.Partitions(topic)
//}

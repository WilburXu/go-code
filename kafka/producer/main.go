package main

import (
	"encoding/json"
	"errors"
	"github.com/Shopify/sarama"
	"log"
	"time"
)

type Client struct {
	sarama.AsyncProducer
	msgPool chan *sarama.ProducerMessage
}

func NewClient() (*Client, error) {
	// create client
	var err error
	c := &Client{
		msgPool: make(chan *sarama.ProducerMessage, 2000),
	}

	config := sarama.NewConfig()
	config.Producer.Return.Errors = true
	config.Producer.Return.Successes = true
	//config.Producer.Partitioner = sarama.NewRandomPartitioner
	//config.Producer.RequiredAcks = sarama.NoResponse
	config.Version = sarama.V2_0_0_0

	c.AsyncProducer, err = sarama.NewAsyncProducer([]string{"127.0.0.1:9092"}, config)
	if err != nil {
		log.Printf("kafka producer connect init err %v \n", err)
		return nil, err
	}

	return c, nil
}

func (c *Client) Run() {
	for {
		select {
		case msg := <-c.msgPool:
			c.Input() <- msg
		}

		select {
		case suc := <-c.Successes():
			log.Printf("message: \"%s\" sent to partition %d at offset %d\n", suc.Value, suc.Partition, suc.Offset)
		case fail := <-c.Errors():
			log.Printf("err: %s\n", fail.Err.Error())
		default:
			log.Printf("ttt %+v \n", c)
			log.Println("kafka producer msg case default...")
		}
	}
}

func (c *Client) Send(topic string, msg []byte) error {
	if topic == "" {
		return errors.New("kafka producer send msg topic empty")
	}

	kafkaMsg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.ByteEncoder(msg),
		Partition: 2,
	}

	c.msgPool <- kafkaMsg

	//c.Input() <- kafkaMsg
	//time.Sleep(time.Second * 1)
	//select {
	//case suc := <-c.Successes():
	//	offset = suc.Offset
	//	log.Printf("message: \"%s\" sent to partition %d at offset %d\n", suc.Value, suc.Partition, suc.Offset)
	//case fail := <-c.Errors():
	//	log.Printf("err: %s\n", fail.Err.Error())
	//	return 0, fail.Err
	//default:
	//	log.Printf("ttt %+v \n", c)
	//	log.Println("kafka producer msg case default...")
	//	return 0, nil
	//}

	return nil
}

func main() {
	kafkaProducer, err := NewClient()
	if err != nil {
		log.Println("222")
	}

	go kafkaProducer.Run()

	i := 1
	for {
		tmp := map[string]interface{}{
			"aa":   i,
			"bb":   "cc",
			"time": time.Now().Format("2006-01-02 15:04:05"),
		}

		log.Println(i)
		bs, _ := json.Marshal(tmp)
		kafkaProducer.Send("web_log", bs)
		time.Sleep(1 * time.Second)
		i++
	}
}

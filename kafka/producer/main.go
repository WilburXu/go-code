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
}

func NewClient() (*Client, error) {
	// create client
	var (
		err error
		c   = new(Client)
	)

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

func (c *Client) Send(topic string, msg []byte) (int64, error) {
	if topic == "" {
		return 0, errors.New("kafka producer send msg topic empty")
	}

	var offset int64

	kafkaMsg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.ByteEncoder(msg),
	}

	c.Input() <- kafkaMsg
	time.Sleep(time.Second * 1)
	select {
	case suc := <-c.Successes():
		offset = suc.Offset
		log.Printf("message: \"%s\" sent to partition %d at offset %d\n", suc.Value, suc.Partition, suc.Offset)
	case fail := <-c.Errors():
		log.Printf("err: %s\n", fail.Err.Error())
		return 0, fail.Err
	default:
		log.Printf("ttt %+v \n", c)
		log.Println("kafka producer msg case default...")
		return 0, nil
	}

	return offset, nil
}

func main() {
	kafkaProducer, err := NewClient()
	if err != nil {
		log.Println("222")
	}

	i:=1
	time.Sleep(time.Second * 2)
	for {
		tmp := map[string]interface{}{
			"aa" : i,
			"bb" : "cc",
		}

		log.Println(i)
		bs, _ := json.Marshal(tmp)
		kafkaProducer.Send("omg-live_down", bs)
		time.Sleep(1*time.Second)
		i++
	}

	//config := sarama.NewConfig()
	//config.Producer.Return.Errors = true
	//config.Producer.Return.Successes = true
	//config.Producer.Partitioner = sarama.NewRandomPartitioner
	//config.Producer.RequiredAcks = sarama.WaitForAll
	//config.Version = sarama.V1_0_0_0
	//
	//producer, _ := sarama.NewAsyncProducer([]string{"127.0.0.1:9092"}, config)
	//
	//log.Println("hongbao kfk broker writer runHongbaoBroker")
	//defer func() {
	//	log.Println("hongbao kfk broker writer stop")
	//	producer.Close()
	//}()
	//
	//for {
	//	log.Println("WilburXu start send msg")
	//
	//	kafkaMsg := &sarama.ProducerMessage{
	//		Topic: "web_log",
	//		Value: sarama.ByteEncoder("WilburXu send msg"),
	//	}
	//	producer.Input() <- kafkaMsg
	//	select {
	//	case suc := <-producer.Successes():
	//		log.Printf("> message: \"%s\" sent to partition  %d at offset %d\n", suc.Value, suc.Partition, suc.Offset)
	//	case fail := <-producer.Errors():
	//		log.Printf("err: %s\n", fail.Err.Error())
	//	default:
	//		log.Println("???? why?")
	//	}
	//
	//	time.Sleep(time.Second * 1)
	//}
}

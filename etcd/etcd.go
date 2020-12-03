package main

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/clientv3/concurrency"
	"log"
	"time"
)

type Client struct {
	*clientv3.Client
}

func NewClient() (*Client, error) {
	//客户端配置
	config := clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	}

	//建立连接
	var (
		err error
		c = new(Client)
	)
	c.Client, err = clientv3.New(config)
	if err != nil {
		log.Println(fmt.Sprintf("client new err: %v", err))
		return nil, err
	}

	timeoutCtx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	_, err = c.Client.Status(timeoutCtx, config.Endpoints[0])
	if err != nil {
		log.Println(fmt.Sprintf("client status err: %v", err))
		return nil, err
	}

	return c, nil
}

func (c *Client) Test() error {
	log.Println("this is etcd client test")
	return nil
}

func (c *Client) SetMutex(etcdCtx context.Context, ctxTimeout int64, lockName string) (*concurrency.Session, *concurrency.Mutex, error) {
	log.Println("is Grant")
	response, err := c.Client.Grant(etcdCtx, ctxTimeout)
	if err != nil {
		log.Fatal(err.Error())
		return nil, nil, err
	}

	log.Println("is NewSession")
	etcdSession, err := concurrency.NewSession(c.Client, concurrency.WithLease(response.ID))
	if err != nil {
		return nil, nil, err
	}

	log.Println("is NewMutex")
	mutexObj := concurrency.NewMutex(etcdSession, lockName)
	if err := mutexObj.Lock(etcdCtx); err != nil {
		log.Fatal(err)
		return nil, nil, err
	}

	return etcdSession, mutexObj, nil
}











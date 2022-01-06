package main

import (
	"fmt"
	"sync"
	"time"
)

type (
	subscriber chan interface{}
	topicFunc  func(v interface{}) bool
)

// 发布者对象
type Publisher struct {
	m           sync.RWMutex
	buffer      int
	timeout     time.Duration
	subscribers map[subscriber]topicFunc // 订阅者信息
}

// 构建一个发布者对象，可以设置发布超时时间和缓存队列的长度
func NewPubisher(publishTimeout time.Duration, buffer int) *Publisher {
	return &Publisher{
		buffer:      buffer,
		timeout:     publishTimeout,
		subscribers: make(map[subscriber]topicFunc),
	}
}

//func (p *Publisher) Subscribe() chan interface{} {
//	return p.SubscribeTopic
//}

func (p *Publisher) SubscribeTopic(topic topicFunc) chan interface{} {
	ch := make(chan interface{}, p.buffer)
	p.m.Lock()
	p.subscribers[ch] = topic
	p.m.Unlock()
	return ch
}

// 发布者对象
func main() {
	fmt.Println("11")
}

package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func main() {
	tr := NewTracker()
	go tr.Run()

	_ = tr.Event(context.Background(), "test1")
	_ = tr.Event(context.Background(), "test2")
	_ = tr.Event(context.Background(), "test3")

	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(7 * time.Second))
	defer cancel()

	tr.Shutdown(ctx)
}

// Tracker init
type Tracker struct {
	ch   chan string
	stop chan struct{}
}

func NewTracker() *Tracker {
	return &Tracker{
		ch: make(chan string, 10),
		stop: make(chan struct{}),
	}
}

func (t *Tracker) Run() {
	// close(t.ch) 后，跳出循环
	for val := range t.ch {
		time.Sleep(time.Second)
		fmt.Println(val)
	}

	// make(chan struct{})，才能生效
	t.stop <- struct{}{}
}

func (t *Tracker) Event(ctx context.Context, data string) error {
	select {
	case t.ch <- data:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

func (t *Tracker) Shutdown(ctx context.Context) {
	defer func(){
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()

	close(t.ch)

	select {
	case <-t.stop:
		log.Println("stop")
	case <-ctx.Done():
		log.Println("done")
	}
}
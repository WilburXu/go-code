package main

import (
	"context"
	"log"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	context.TODO()
	defer cancel()

	go handle(ctx, time.Millisecond * 500)

	select {

	}
}

func handle(ctx context.Context, duration time.Duration) {
	select {
	case <-ctx.Done():
		log.Println("done", ctx.Err())
	case <-time.After(duration):
		log.Println("after", duration)
	}
}

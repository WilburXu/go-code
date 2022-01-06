package c_t4

import (
	"context"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second)
}
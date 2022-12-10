package limiter

import (
	"fmt"
	"testing"
	"time"
)

type TokenBucketLimiter struct {
	tokens chan struct{}
	close  chan struct{}
}

func TestLimiter(t *testing.T) {
	buffer, interval := 0, 500*time.Millisecond
	res := &TokenBucketLimiter{
		tokens: make(chan struct{}, buffer),
		close:  make(chan struct{}),
	}
	go func() {
		producer := time.NewTicker(interval)
		defer producer.Stop()
		for t := range producer.C {
			fmt.Println(t)
			res.tokens <- struct{}{}
		}
	}()
	time.Sleep(13 * time.Second)
	fmt.Println("sleep end...")
	i := 0
	for {
		i++
		fmt.Println(<-res.tokens, i)
	}
}

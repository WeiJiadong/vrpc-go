package queue

import (
	"context"
	"sync"
	"testing"
	"time"

	"gopkg.in/go-playground/assert.v1"
)

func TestQueue_Put(t *testing.T) {
	q := NewQueue(1)
	wg := sync.WaitGroup{}
	for i := 0; i < 100000; i++ {
		i := i
		// 生产者
		wg.Add(1)
		go func() {
			defer wg.Done()
			ctx, _ := context.WithTimeout(context.TODO(), time.Second)
			assert.Equal(t, q.Put(ctx, i), nil)
		}()
		// 消费者
		wg.Add(1)
		go func() {
			defer wg.Done()
			ctx, _ := context.WithTimeout(context.TODO(), time.Second)
			q.Get(ctx)
		}()
	}

	wg.Wait()
}

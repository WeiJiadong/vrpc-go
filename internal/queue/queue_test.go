package queue

import (
	"sync"
	"testing"
	"time"
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
			q.Put(i, time.Second)
		}()
		// 消费者
		wg.Add(1)
		go func() {
			defer wg.Done()
			q.Get(time.Second)
		}()
	}

	wg.Wait()
}

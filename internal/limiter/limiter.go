package limiter

import (
	"time"
)

type Limiter struct {
	token chan struct{}
}

func NewLimiter(interval time.Duration) {
	res := &Limiter{
		token: make(chan struct{}),
	}
	go func() {
		producer := time.NewTicker(interval)
		defer producer.Stop()
		for range producer.C {
			res.token <- struct{}{}
		}
	}()
}

func (l *Limiter) Allow() bool {
	<-l.token
	return true
}

func (l *Limiter) Close() {
	close(l.token)
}

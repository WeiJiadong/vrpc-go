package queue

import (
	"errors"
	"sync"
	"time"
)

// Queue 定义队列的结构体
type Queue struct {
	// 存储队列中的数据
	data   []any
	length int

	// 使用互斥锁保证线程安全
	mu     *sync.Mutex
	dataMu *sync.Mutex
	// 使用条件变量来通知队列的状态变化
	condW *sync.Cond
	condR *sync.Cond
}

// 初始化队列
// 参数 n 表示队列的最大容量
func NewQueue(n int) *Queue {
	q := &Queue{
		// 初始化队列
		data: make([]any, 0, n),
		// 初始化互斥锁
		mu:       &sync.Mutex{},
		dataMu: &sync.Mutex{},
	}
	// 初始化条件变量
	q.condW = sync.NewCond(q.mu)
	q.condR = sync.NewCond(q.mu)

	q.length = n

	return q
}

// 向队列中插入一个元素
// 参数 t 表示超时时间
func (q *Queue) Put(v any, t time.Duration) error {
	// 加锁
	q.mu.Lock()
	defer q.mu.Unlock()

	// 使用 select 语句来实现超时处理
	select {
	// 如果在 t 时间内获取到锁，则执行后续操作
	case <-time.After(t):
		return errors.New("timeout")
	default:
		// 检查队列是否已满
		for len(q.data) == q.length {
			// 等待队列有空闲位置
			q.condR.Wait()
		}

		// 向队列中插入元素
		q.data = append(q.data, v)

		// 通知队列可写
		q.condW.Signal()

		return nil
	}
}

// 从队列中取出一个元素
// 参数 t 表示超时时间
func (q *Queue) Get(t time.Duration) (any, error) {
	// 加锁
	q.mu.Lock()
	defer q.mu.Unlock()

	// 使用 select 语句来实现超时处理
	select {
	// 如果在 t 时间内获取到锁，则执行后续操作
	case <-time.After(t):
		return 0, errors.New("timeout")
	default:
		// 检查队列是否为空
		for len(q.data) == 0 {
			// 等待队列有元素
			q.condW.Wait()
		}

		// 从队列中取出元素
		v := q.data[0]
		q.data = q.data[1:]

		// 通知队列可读了
		q.condR.Signal()

		return v, nil
	}
}

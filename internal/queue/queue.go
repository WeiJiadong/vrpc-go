package queue

import (
	"context"
	"sync"
)

// Queue 定义队列的结构体
type Queue struct {
	// 存储队列中的数据
	data   []any
	length int

	// 使用互斥锁保证线程安全
	mu *sync.Mutex
	// 使用条件变量来通知队列的状态变化
	canWrite *sync.Cond
	canRead  *sync.Cond
}

// 初始化队列
// 参数 n 表示队列的最大容量
func NewQueue(n int) *Queue {
	q := &Queue{
		// 初始化队列
		data: make([]any, 0, n),
		// 初始化互斥锁
		mu: &sync.Mutex{},
	}
	// 初始化条件变量
	q.canWrite = sync.NewCond(q.mu)
	q.canRead = sync.NewCond(q.mu)

	q.length = n

	return q
}

// 向队列中插入一个元素
// 参数 t 表示超时时间
func (q *Queue) Put(ctx context.Context, v any) error {
	// 加锁
	q.mu.Lock()
	defer q.mu.Unlock()

	// 超时检测
	if ctx.Err() != nil {
		return ctx.Err()
	}

	// 检查队列是否已满
	for len(q.data) == q.length {
		// 等待队列可写
		q.canWrite.Wait()
		if ctx.Err() != nil {
			return ctx.Err()
		}
	}

	// 向队列中插入元素
	q.data = append(q.data, v)

	// 通知队列可读
	q.canRead.Signal()

	return nil

}

// 从队列中取出一个元素
// 参数 t 表示超时时间
func (q *Queue) Get(ctx context.Context) (any, error) {
	// 加锁
	q.mu.Lock()
	defer q.mu.Unlock()

	// 超时检测
	if ctx.Err() != nil {
		return nil, ctx.Err()
	}

	// 检查队列是否为空
	for len(q.data) == 0 {
		// 等待队列可读
		q.canRead.Wait()
		if ctx.Err() != nil {
			return nil, ctx.Err()
		}
	}

	// 从队列中取出元素
	v := q.data[0]
	q.data = q.data[1:]

	// 通知队列可写了
	q.canWrite.Signal()

	return v, nil
}

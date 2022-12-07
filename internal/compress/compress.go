// Package compress 压缩算法实现
package compress

import "sync"

// Compressor 压缩功能实现
type Compressor interface {
	Compress(data []byte) ([]byte, error)
	Uncompress(data []byte) ([]byte, error)
}

// Compress 压缩算法管理结构
type Compress struct {
	coms map[byte]Compressor
	mtx  sync.Mutex
}

// Register 注册压缩算法实现
func (c *Compress) Register(num byte, com Compressor) {
	c.mtx.Lock()
	defer c.mtx.Unlock()
	c.coms[num] = com
}

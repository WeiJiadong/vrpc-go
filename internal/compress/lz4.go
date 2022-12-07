package compress

import (
	"github.com/pierrec/lz4/v4"
)

// Lz4Compress Lz4 实现
type Lz4Compress struct {
	lz4.Compressor
	size int
}

// Compress Lz4压缩
func (lc *Lz4Compress) Compress(data []byte) ([]byte, error) {
	buf := make([]byte, lz4.CompressBlockBound(len(data)))
	n, err := lc.CompressBlock(data, buf)
	if err != nil {
		return nil, err
	}
	lc.size = len(data)
	return buf[:n], nil
}

// Uncompress Lz4 解压缩
func (lc *Lz4Compress) Uncompress(data []byte) ([]byte, error) {
	val := make([]byte, lc.size)
	n, err := lz4.UncompressBlock(data, val)
	if err != nil {
		return nil, err
	}
	return val[:n], nil
}

package compress

import (
	"bytes"
	"compress/gzip"
	"io/ioutil"
)

// GzipCompress gzip 实现
type GzipCompress struct {
}

// Compress gzip压缩
func (gc *GzipCompress) Compress(data []byte) ([]byte, error) {
	buf := bytes.Buffer{}
	w := gzip.NewWriter(&buf)
	if _, err := w.Write(data); err != nil {
		return nil, err
	}
	if err := w.Close(); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// Uncompress gzip 解压缩
func (gc *GzipCompress) Uncompress(data []byte) ([]byte, error) {
	buf := bytes.NewBuffer(data)
	r, err := gzip.NewReader(buf)
	if err != nil {
		return nil, err
	}
	val, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	return val, nil
}

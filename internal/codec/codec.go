// Package codec 包含协议打解包
package codec

import (
	"encoding/binary"
	"encoding/json"
	"unsafe"
)

// Marshaler 序列化操作
type Marshaler func() ([]byte, error)

// Unmarshaler 反序列化操作
type Unmarshaler func([]byte) error

// Message 消息接口定义
type Message interface {
	Marshaler
	Unmarshaler
}

// VrpcMsg vrpc 消息结构定义
type VrpcMsg struct {
	MsgLen              uint32            // VrpcMsg 包大小
	BodyLen             uint32            // 数据Body 大小
	MsgID               int64             // 消息id
	Vesion              byte              // 协议版本
	Compress            byte              // 压缩算法
	Serialization       byte              // Body序列化方式
	CalleeServerNameLen byte              // 被调服务名长度
	CalleeMethodNameLen byte              // 被调方法名
	CalleeServerName    string            // 被调服务名
	CalleeMethodName    string            // 被调方法名
	Mate                map[string]string // 元数据
	Body                []byte            // 实际数据
}

// Marshal 将VrpcMsg序列化为字节数组
func (c *VrpcMsg) Marshal() ([]byte, error) {
	// 1 先计算mate序列化后的长度，与MsgLen的值相关
	val, err := json.Marshal(c.Mate)
	if err != nil {
		return nil, err
	}
	// 2 再计算MsgLen
	size := 4 + 4 + 8 + 1 + 1 + 1 + 1 + 1
	size += int(c.CalleeServerNameLen)
	size += int(c.CalleeMethodNameLen)
	size += len(val)
	size += len(c.Body)
	c.MsgLen = uint32(size)
	// 3 序列化字段
	data := make([]byte, 0, unsafe.Sizeof(*c))
	tmp := make([]byte, unsafe.Sizeof(*c))
	binary.BigEndian.PutUint32(tmp, c.MsgLen)
	data = append(data, tmp[:unsafe.Sizeof(c.MsgLen)]...)
	binary.BigEndian.PutUint32(tmp, c.BodyLen)
	data = append(data, tmp[:unsafe.Sizeof(c.BodyLen)]...)
	binary.BigEndian.PutUint64(tmp, uint64(c.MsgID))
	data = append(data, tmp[:unsafe.Sizeof(c.MsgID)]...)
	data = append(data, c.Vesion)
	data = append(data, c.Compress)
	data = append(data, c.Serialization)
	data = append(data, c.CalleeServerNameLen)
	data = append(data, c.CalleeMethodNameLen)
	data = append(data, []byte(c.CalleeServerName)...)
	data = append(data, []byte(c.CalleeMethodName)...)
	data = append(data, val...)
	data = append(data, c.Body...)
	return data, nil
}

// Unmarshal将字节数组反序列化为VrpcMsg结构
func (c *VrpcMsg) Unmarshal(data []byte) error {
	c.MsgLen = binary.BigEndian.Uint32(data)
	cur := int(unsafe.Sizeof(c.MsgLen))
	c.BodyLen = binary.BigEndian.Uint32(data[cur:])
	cur += int(unsafe.Sizeof(c.BodyLen))
	c.MsgID = int64(binary.BigEndian.Uint64(data[cur:]))
	cur += int(unsafe.Sizeof(c.MsgID))
	c.Vesion = data[cur]
	cur++
	c.Compress = data[cur]
	cur++
	c.Serialization = data[cur]
	cur++
	c.CalleeServerNameLen = data[cur]
	cur++
	c.CalleeMethodNameLen = data[cur]
	cur++
	c.CalleeServerName = string(data[cur : cur+int(c.CalleeServerNameLen)])
	cur += int(c.CalleeServerNameLen)
	c.CalleeMethodName = string(data[cur : cur+int(c.CalleeMethodNameLen)])
	cur += int(c.CalleeMethodNameLen)
	err := json.Unmarshal(data[cur:c.MsgLen-c.BodyLen], &c.Mate)
	if err != nil {
		return err
	}
	cur = int(c.MsgLen - c.BodyLen)
	c.Body = data[cur:]
	return nil
}

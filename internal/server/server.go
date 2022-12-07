// Package server rpc Server各种协议的实现
package server

// Server rpc server 结构定义
type Server interface {
	// Serve 启动服务
	Serve() error
	// Close 关闭服务
	Close() error
}

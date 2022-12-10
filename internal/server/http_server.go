package server

import (
	"net/http"
	"sync"

	"github.com/WeiJiadong/vrpc-go/internal/compress"
	"github.com/WeiJiadong/vrpc-go/internal/service"
)

// HttpServer http server定义
type HttpServer struct {
	opt *HttpServerOpt
	http.Server
}

// HttpServerOpt http server选项参数
type HttpServerOpt struct {
	services map[string]service.Service
	mtx      sync.Mutex
	addr     string
	compress compress.Compress
}

// HttpServerOptHelper HttpServerOpt helper
type HttpServerOptHelper func(hso *HttpServerOpt)

// Register 注册service
func (hs *HttpServer) Register(name string, fn service.Service) {
	hs.opt.mtx.Lock()
	defer hs.opt.mtx.Unlock()
	hs.opt.services[name] = fn
}

// Serve 启动服务
func (hs *HttpServer) Serve() error {
	return nil
}

// Close 关闭服务
func (hs *HttpServer) Close() error {
	return nil
}

// NewHttpServer http server构造函数
func NewHttpServer(addr string, opts ...HttpServerOptHelper) *HttpServer {
	opt := &HttpServerOpt{
		addr:     addr,
		services: make(map[string]service.Service),
	}
	for i := range opts {
		opts[i](opt)
	}
	return &HttpServer{
		opt: opt,
	}
}

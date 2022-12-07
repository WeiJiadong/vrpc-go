// package service rpc service结构定义
package service

import "context"

// Fn 方法对应的函数签名
type Fn func(ctx context.Context, req any) (any, error)

// Method service支持的方法集合定义
type Method struct {
	Fns map[string]Fn
}

// Service service 结构定义
type Service struct {
	route map[string][]Method
}

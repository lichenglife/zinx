package znet

import "lichenglife/zinx/ziface"

// 定义基础路由
// 通过不同路由类型实现基础路由，实现PreHandler、Handler 、PostHandler 的自由组合
type BaseRouter struct {
}

// 前置函数
func (r *BaseRouter) PreHandler(request ziface.IRequest) {

}

// 业务函数
func (r *BaseRouter) Handler(request ziface.IRequest) {

}

// 后置函数
func (r *BaseRouter) PostHandler(request ziface.IRequest) {

}

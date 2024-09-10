package main

import (
	"fmt"
	"lichenglife/zinx/ziface"
	"lichenglife/zinx/znet"
)

func main() {

	// TODO 定义Router

	fmt.Println("Start  Zinx Server v0.1")
	//1、 创建Server
	s := znet.NewServer()
	// 2、注册路由函数
	p := &PingRouter{}

	s.AddRouter(p)
	// 3、启动服务
	s.Serve()

}

// 接口组合实现多态
type PingRouter struct {
	znet.BaseRouter
}

func (p *PingRouter) PreHandler(request ziface.IRequest) {
	connection := request.GetConn()

	id := connection.GetConnID()

	fmt.Printf("current connection id is %d \n", id)
}
func (p *PingRouter) Handler(request ziface.IRequest) {

	fmt.Println("Finish to exutate the  Handler")
}

// 定义PostHandler
func (p *PingRouter) PostHandler(request ziface.IRequest) {

	fmt.Println("Start exutate PostHandler........")

	data := make([]byte, 512)
	_, err := request.GetConn().GetConnection().Read(data)
	if err != nil {
		fmt.Println("PostHandler error", err)
		return
	}
	fmt.Println("Finish to exutate the post Handler")

}

//  server 注册Router

//  Router 通过request 执行 handler

//  Router 管理Connection  将Connection与handler 进行绑定

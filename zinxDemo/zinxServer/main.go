package main

import (
	"fmt"
	"lichenglife/zinx/ziface"
	"lichenglife/zinx/znet"
)

func main() {

	// TODO 定义Router

	fmt.Println("Start  Zinx Server v0.5")
	//1、 创建Server
	s := znet.NewServer()
	// 2、注册路由函数
	p := &PingRouter{}

	h := &HelloRouter{}

	m := &MessageRouter{}

	t := &ThreeRouter{}

	s.AddRouter(0, p)
	s.AddRouter(1, h)
	s.AddRouter(2, m)
	s.AddRouter(3, t)
	// 3、启动服务
	s.Serve()

}

type ThreeRouter struct {
	znet.BaseRouter
}

func (h *ThreeRouter) Handler(request ziface.IRequest) {
	fmt.Println("Call MessageRouter MessageHandler")
	// 先读取客户端的数据，再回写hello...hello...hello
	fmt.Println("recv from client : msgId=", request.GetMsgID(), ", data=", string(request.GetData()))

	// 回写
	err := request.GetConn().SendMsg(2, []byte("three...three...three"))
	if err != nil {
		fmt.Println(err)
	}
}

type MessageRouter struct {
	znet.BaseRouter
}

func (h *MessageRouter) Handler(request ziface.IRequest) {
	fmt.Println("Call MessageRouter MessageHandler")
	// 先读取客户端的数据，再回写hello...hello...hello
	fmt.Println("recv from client : msgId=", request.GetMsgID(), ", data=", string(request.GetData()))

	// 回写
	err := request.GetConn().SendMsg(2, []byte("message...message...message"))
	if err != nil {
		fmt.Println(err)
	}
}

type HelloRouter struct {
	znet.BaseRouter
}

func (h *HelloRouter) PreHandler(request ziface.IRequest) {
	fmt.Println("Call PreRouter PreHandler")
	// 先读取客户端的数据，再回写hello...hello...hello
	fmt.Println("recv from client : msgId=", request.GetMsgID(), ", data=", string(request.GetData()))

	// 回写
	err := request.GetConn().SendMsg(1, []byte("hello...hello...hello"))
	if err != nil {
		fmt.Println(err)
	}
}
func (h *HelloRouter) Handler(request ziface.IRequest) {
	fmt.Println("Call HelloRouter HelloHandle")
	// 先读取客户端的数据，再回写hello...hello...hello
	fmt.Println("recv from client : msgId=", request.GetMsgID(), ", data=", string(request.GetData()))

	// 回写
	err := request.GetConn().SendMsg(1, []byte("hello...hello...hello"))
	if err != nil {
		fmt.Println(err)
	}
}

// 接口组合实现多态
type PingRouter struct {
	znet.BaseRouter
}

func (p *PingRouter) PreHandler(request ziface.IRequest) {
	fmt.Println("Call PingRouter PreHandle")
	connection := request.GetConn()

	id := connection.GetConnID()

	fmt.Printf("current connection id is %d \n", id)
}
func (p *PingRouter) Handler(request ziface.IRequest) {

	fmt.Println("Call PingRouter Handle")
	//先读取客户端的数据，再回写ping...ping...ping
	fmt.Println("recv from client : msgId=", request.GetMsgID(), ", data=", string(request.GetData()))

	//回写
	err := request.GetConn().SendMsg(0, []byte("ping...ping...ping"))
	if err != nil {
		fmt.Println(err)
	}
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

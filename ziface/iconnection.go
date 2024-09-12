package ziface

import "net"

/*
对于服务端在接收到不同的客户端， 需要进行处理不同的业务函数

实现上是通过 将不同的客户端连接 绑定到不同的业务函数上

所以第一步将客户端连接进行抽象处理
*/
type IConnection interface {

	// 启动当前连接
	Start()
	// 停止当前连接

	Stop()

	//  获取连接ID

	GetConnID() uint32

	//  获取当前连接

	GetConnection() *net.TCPConn

	// 发送消息  无缓冲

	SendMsg(msgId uint32, data []byte) error

	SendBuffMsg(msgId uint32, data []byte) error

	// 增加连接参数
	SetProperty(key string, value interface{})

	GetProperty(key string) (interface{}, error)

	// 删除
	RemoveProperty(key string)
}

// 定义一个统一处理业务函数的函数接口类型, 后续业务处理函数都属于该类型
// 通过路由将 连接与 业务函数进行绑定
type HandFunc func(*net.TCPConn, []byte, int) error

//  业务抽象
//  业务建模

//  Request
//  Message
//  Connection
//  Router
//  handler
//  Server

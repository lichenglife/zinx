package znet

import (
	"fmt"
	"lichenglife/zinx/utils"
	"lichenglife/zinx/ziface"
	"net"
)

// 服务器模块

type Server struct {

	//  服务器名称
	Name string
	//  监听端口
	Port int
	//  监听地址
	IP string

	//  服务器版本
	IPVersion string

	// 当前服务的MsgHandler模块，用于管理 消息以及对应的处理函数
	MsgHandler ziface.IMsgHandler
}

// 实例化Server 对象
func NewServer() ziface.IServer {

	utils.GlobalObject.Reload()

	server := &Server{
		Name:       utils.GlobalObject.Name,
		IPVersion:  "tcp4",
		IP:         utils.GlobalObject.Host,
		Port:       utils.GlobalObject.TcpPort,
		MsgHandler: NewMsgHandler(),
	}
	return server
}

// 启动网络服务
func (s *Server) Start() {

	//   监听端口
	fmt.Printf("Start Server  address %s, port %d , name %s \n", s.IP, s.Port, s.Name)

	//  开启协程提供处理客户端请求

	go func() {

		// 0、启动WorkPool
		s.MsgHandler.StartWorkPool()

		//  1、获取TCP的地址
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("ResolveTCPAddr error", err)
			return
		}

		//  2、监听服务器的地址
		listener, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Println("Listen error", err)
			return
		}

		// 3、启动网络监听，处理客户端请求数据
		for {
			// 阻塞等待客户端连接
			conn, err := listener.AcceptTCP()

			//  封装Request
			if err != nil {
				fmt.Println("Accept error", err)
				continue
			}
			var connID uint32
			c := NewConnection(conn, connID, s.MsgHandler)
			connID++
			go c.Start()

		}

	}()

}

// 停止服务
func (s *Server) Stop() {
	fmt.Println("[STOP] Zinx server , name ", s.Name)
}

// 运行
func (s *Server) Serve() {

	s.Start()

	//  防止主进程退出, 进行阻塞
	select {}

}

// 添加路由
func (s *Server) AddRouter(msgID uint32, router ziface.IRouter) {

	s.MsgHandler.AddRouter(msgID, router)

}

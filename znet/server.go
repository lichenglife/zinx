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

	// 路由函数， 函数数组
	Router ziface.IRouter
}

// 定义处理函数
// func dealhandler(conn *net.TCPConn, data []byte,) error {
// 	cnt, err := conn.Read(data)
// 	if err != nil {
// 		fmt.Println("get data from client errror", err)
// 		return errors.New("get data from client errror")
// 	}
// 	fmt.Printf(",get data from client,%s,%d \n ", string(data), cnt)
// 	fmt.Println("current handler dealHandler")

// 	//  通过匿名函数实现业务函数调用

// 	//  TODO  定义函数，通过路由实现绑定函数
// 	_, err = conn.Write(data[:cnt])
// 	if err != nil {
// 		fmt.Println("send data to client errror", err)
// 		return errors.New("get data from client errror")
// 	}

// 	return nil
// }

// 启动网络服务
func (s *Server) Start() {

	//   监听端口
	fmt.Printf("Start Server  address %s, port %d , name %s \n", s.IP, s.Port, s.Name)

	//  开启协程提供处理客户端请求

	go func() {

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
			c := NewConnection(conn, connID, s.Router)
			connID++
			go c.Start()

		}

	}()

}

//   停止服务

func (s *Server) Stop() {
	fmt.Println("[STOP] Zinx server , name ", s.Name)
}

// 运行
func (s *Server) Serve() {

	s.Start()

	//  防止主进程退出, 进行阻塞
	select {}

}

func (s *Server) AddRouter(router ziface.IRouter) {
	s.Router = router
}

// 实例化Server 对象
func NewServer() ziface.IServer {

	utils.GlobalObject.Reload()

	server := &Server{
		Name:      utils.GlobalObject.Name,
		IPVersion: "tcp4",
		IP:        utils.GlobalObject.Host,
		Port:      utils.GlobalObject.TcpPort,
		Router:    nil,
	}
	return server
}

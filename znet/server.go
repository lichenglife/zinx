package znet

import (
	"fmt"
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
}

// 启动网络服务
func (s *Server) Start() {

	//   监听端口
	fmt.Printf("Start Server  address %s, port %d \n", s.IP, s.Port)

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
			if err != nil {
				fmt.Println("Accept error", err)
				continue
			}

			//  并发获取客户端请求数据
			go func() {

				data := make([]byte, 512)

				for {
					cnt, err := conn.Read(data)
					if err != nil {
						fmt.Println("get data from client errror", err)
						continue
					}
					fmt.Printf("get data from client %s,%d ", string(data), cnt)

					//  回写数据
					_, err = conn.Write(data[:cnt])
					if err != nil {
						fmt.Printf("send data to client errror", err)
						continue
					}

				}

			}()

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

// 实例化Server 对象
func NewServer(name string, ip string, port int) ziface.IServer {

	server := &Server{
		Name:      "ZinxServerApp V0.1",
		IPVersion: "tcp4",
		IP:        "127.0.0.1",
		Port:      8999,
	}
	return server
}

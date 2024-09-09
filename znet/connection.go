package znet

import (
	"fmt"
	"lichenglife/zinx/ziface"
	"net"
)

type Connection struct {

	// 当前连接的套接字
	Conn *net.TCPConn

	//  连接ID
	ConnID uint32

	// 当前连接的状态
	IsClosed bool

	// 绑定连接处理的函数
	handlerApi ziface.HandFunc

	// 定义无缓冲通道，用于记录该连接的退出状态
	ExitBuffChan chan bool
}

// 启动当前连接
func (c *Connection) Start() {

	// TODO  执行绑定处理的业务函数

	fmt.Printf("current ConnID is %d Starting \n", c.ConnID)

	// 执行读函数

	// 执行写函数

	// for {
	// 	// start to read from  client request
	// 	data := make([]byte, 512)

	// 	_, err := c.Conn.Read(data)
	// 	if err != nil {
	// 		fmt.Printf("read from client request error ,current ConnID is %d", c.GetConnID())
	// 		return
	// 	}
	// 	fmt.Printf("Current ConnID: %d,Read from Request: %s, \n", c.ConnID, string(data))

	// }

	data := make([]byte, 512)
	err := c.handlerApi(c.Conn, data, len(data))
	if err != nil {

		return
	}

	for {

		select {

		case <-c.ExitBuffChan:
			return
		}
	}

}

// 停止当前连接

func (c *Connection) Stop() {

	if c.IsClosed {
		return
	}

	//todo  关闭连接

	c.ExitBuffChan <- true

	close(c.ExitBuffChan)

}

//  获取连接ID

func (c *Connection) GetConnID() uint32 {

	return c.ConnID
}

// 实例化连接
func NewConnection(conn *net.TCPConn, connID uint32, handlerApi ziface.HandFunc) ziface.IConnection {

	c := &Connection{
		Conn:         conn,
		ConnID:       connID,
		handlerApi:   handlerApi,
		IsClosed:     false,
		ExitBuffChan: make(chan bool, 1),
	}

	return c

}

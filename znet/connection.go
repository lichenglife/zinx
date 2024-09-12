package znet

import (
	"errors"
	"fmt"
	"io"
	"lichenglife/zinx/utils"
	"lichenglife/zinx/ziface"
	"net"
	"sync"
)

type Connection struct {

	//  结构体组合
	// 1、分布式架构部署，可标识当前连接属于哪个Server
	// 2、接口组合，使用Server的方法函数
	TcpSer ziface.IServer

	// 当前连接的套接字
	Conn *net.TCPConn

	//  连接ID
	ConnID uint32

	// 当前连接的状态
	IsClosed bool

	// 绑定连接处理的函数
	//handlerApi ziface.HandFunc

	// server 注册Router
	// MsgHandler 注册 msgId Router
	// connection绑定 Msghandler
	// Msghandler 执行DoHandler 执行Router函数
	MsgHandler ziface.IMsgHandler

	// 定义无缓冲通道，用于记录该连接的退出状态
	ExitBuffChan chan bool

	// 定义无缓冲通道， 用于读写两个goroutine 之间进行消息通信
	msgChan chan []byte

	// 定义有缓存通道， 多个读写Goroutine  避免阻塞
	msgBuffChan chan []byte

	//  连接属性
	Property map[string]interface{}

	//保护链接属性修改的锁 map操作并发操作
	propertyLock sync.RWMutex
}

// 增加连接参数
func (c *Connection) SetProperty(key string, value interface{}) {
	c.propertyLock.Lock()
	defer c.propertyLock.Unlock()
	if _, ok := c.Property[key]; ok {
		c.Property[key] = value
	}
}

func (c *Connection) GetProperty(key string) (interface{}, error) {

	c.propertyLock.Lock()
	defer c.propertyLock.Unlock()

	if value, ok := c.Property[key]; ok {
		return value, nil
	}
	return nil, errors.New("key is not exist")
}

// 删除
func (c *Connection) RemoveProperty(key string) {

	c.propertyLock.Lock()
	defer c.propertyLock.Unlock()
	delete(c.Property, key)
}

// 启动写 Goroutine
func (c *Connection) StartWriter() {

	fmt.Println("Start Writer Goroutine")

	defer fmt.Println(c.Conn.RemoteAddr(), "Close")

	for {

		select {
		case data := <-c.msgChan:
			// 通道有数据写回客户端
			if _, err := c.Conn.Write(data); err != nil {
				fmt.Println("Writer Goroutine writer data error", err)
				return
			}
		case data, ok := <-c.msgBuffChan:
			if ok {
				if _, err := c.Conn.Write(data); err != nil {
					fmt.Println("Writer Goroutine writer data error", err)
					return
				} else {
					fmt.Println("MsgBuffChan is Closed")
					break
				}
			}

		case <-c.ExitBuffChan:
			// Read Goroutine notice connection exit
			return
		}
	}

}
func (c *Connection) StartReader() {

	fmt.Println("[Reader Goroutine is running]")
	defer fmt.Println(c.GetConnID(), "[conn Reader exit!]")
	defer c.Stop()

	for {

		// 读取客户端数据
		// 创建拆解包对象
		dp := NewDataPack()
		// 读取客户端请求的 msgHead
		headData := make([]byte, dp.GetHeadLen())
		if _, err := io.ReadFull(c.GetConnection(), headData); err != nil {
			fmt.Println("read msg head error", err)
			c.ExitBuffChan <- true
			continue
		}
		// 根据msgHeadData 获得msgID、dataLen
		msg, err := dp.UpPack(headData)
		if err != nil {
			fmt.Printf("unpack error %v", err)
			c.ExitBuffChan <- true
			continue
		}
		// 根据dataLen  读取data ,放在msg.data
		var data []byte
		if msg.GetDataLen() > 0 {
			data = make([]byte, msg.GetDataLen())
			if _, err := io.ReadFull(c.GetConnection(), data); err != nil {
				fmt.Println("read  msg data error", err)

				c.ExitBuffChan <- true
				continue
			}

		}
		msg.SetData(data)

		//  这里不是可以通过拆包将请求数据包直接读取出来， 为什么先拆包读取
		r := Request{conn: c, message: msg}

		if utils.GlobalObject.WorkPoolSize > 0 {
			// 将客户端请求 添加到工作队列中
			c.MsgHandler.SendMsgToTaskQueue(&r)
		} else {
			go c.MsgHandler.DoMsgHandler(&r)
		}

	}
}

// 启动当前连接
func (c *Connection) Start() {

	// TODO  执行绑定处理的业务函数

	fmt.Printf("current ConnID is %d Starting \n", c.ConnID)

	go c.StartReader()

	go c.StartWriter()

	// 启动CallOnConnStart
	c.TcpSer.CallOnConnStart(c)

	// 执行读函数

	// 执行写函数
	for {
		select {

		case <-c.ExitBuffChan:
			return
		}
	}

}

// 停止当前连接

func (c *Connection) Stop() {

	c.TcpSer.CallOnConnStop(c)
	if c.IsClosed {
		return
	}
	c.IsClosed = true
	// 在销毁连接之前，停止连接

	c.TcpSer.GetConnMgr().Remove(c.ConnID)
	//todo  关闭连接

	c.ExitBuffChan <- true

	close(c.ExitBuffChan)
	close(c.msgBuffChan)

}

//  获取连接ID

func (c *Connection) GetConnID() uint32 {

	return c.ConnID
}

func (c *Connection) GetConnection() *net.TCPConn {

	return c.Conn
}

// handler sendMsg 发送数据到客户端
func (c *Connection) SendMsg(msgId uint32, data []byte) error {

	if c.IsClosed {
		return errors.New("connection is closed when send msg")
	}
	// 封包
	dp := NewDataPack()

	msgData, err := dp.Pack(NewMsgPackage(data, msgId))
	if err != nil {
		fmt.Println("pack message error", err)
		c.ExitBuffChan <- true
		return errors.New("pack message error")
	}

	// 修改为通过WriterGoroutine 写数据
	c.msgChan <- msgData
	// if _, err := c.GetConnection().Write(msgData); err != nil {
	// 	fmt.Println("writer data to client error", err)
	// 	c.ExitBuffChan <- true
	// 	return errors.New("writer data to client error")
	// }
	return nil
}

func (c *Connection) SendBuffMsg(msgId uint32, data []byte) error {

	if c.IsClosed {
		return errors.New("connection is closed when send msg")
	}
	// 封包
	dp := NewDataPack()

	msgData, err := dp.Pack(NewMsgPackage(data, msgId))
	if err != nil {
		fmt.Println("pack message error", err)
		c.ExitBuffChan <- true
		return errors.New("pack message error")
	}

	// 修改为通过WriterGoroutine 写数据
	c.msgBuffChan <- msgData
	// if _, err := c.GetConnection().Write(msgData); err != nil {
	// 	fmt.Println("writer data to client error", err)
	// 	c.ExitBuffChan <- true
	// 	return errors.New("writer data to client error")
	// }
	return nil

}

// 实例化连接
func NewConnection(tcpSer ziface.IServer, conn *net.TCPConn, connID uint32, mh ziface.IMsgHandler) ziface.IConnection {

	c := &Connection{
		TcpSer:       tcpSer,
		Conn:         conn,
		ConnID:       connID,
		MsgHandler:   mh,
		IsClosed:     false,
		ExitBuffChan: make(chan bool, 1),
		msgChan:      make(chan []byte),
		// 有缓存通道
		msgBuffChan: make(chan []byte, utils.GlobalObject.MaxMsgChanLen),
		Property:    make(map[string]interface{}),
	}

	// 创建连接后添加到ConnMgr中
	c.TcpSer.GetConnMgr().Add(c)
	return c

}

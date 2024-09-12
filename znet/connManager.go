package znet

import (
	"errors"
	"fmt"
	"lichenglife/zinx/utils"
	"lichenglife/zinx/ziface"
	"sync"
)

type ConnManager struct {

	// 连接集合
	Conns map[uint32]ziface.IConnection

	// 读写锁 开箱即用
	connLock sync.RWMutex
}

func NewConnManager() ziface.IConnManager {

	conn := &ConnManager{

		Conns: make(map[uint32]ziface.IConnection),
	}

	return conn
}

// 添加连接
func (c *ConnManager) Add(conn ziface.IConnection) {
	c.connLock.Lock()
	defer c.connLock.Unlock()

	if utils.GlobalObject.MaxConn > 0 {

		c.Conns[conn.GetConnID()] = conn

	}

}

// 删除连接
func (c *ConnManager) Remove(connID uint32) {
	c.connLock.Lock()
	defer c.connLock.Unlock()

	delete(c.Conns, connID)

	fmt.Printf("删除连接 connID %d \n", connID)

}

// 获取连接
func (c *ConnManager) GetConn(connID uint32) (conn ziface.IConnection, err error) {
	c.connLock.Lock()
	defer c.connLock.Unlock()

	if conn, ok := c.Conns[connID]; ok {

		return conn, nil
	}

	return nil, errors.New("连接不存在" + string(connID))
}

// 获取当前连接池大小
func (c *ConnManager) Len() int {
	fmt.Printf("当前连接池大小为: %d \n", len(c.Conns))
	return len(c.Conns)
}

// 清空连接池
func (c *ConnManager) ClearConn() {
	c.connLock.Lock()
	defer c.connLock.Unlock()

	conns := c.Conns
	if len(conns) > 0 {
		//  遍历删除连接
		for connID, conn := range conns {
			//  停止当前连接 通知使用方连接退出、 释放连接
			conn.Stop()
			delete(conns, connID)
		}
	}

	fmt.Println("Clear All Connections successfully: conn num = ", c.Len())
}

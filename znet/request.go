package znet

import "lichenglife/zinx/ziface"

// 定义Request对象, 抽象连接以及请求数据
type Request struct {
	conn ziface.IConnection
	data []byte
}

func (r *Request) GetConn() ziface.IConnection {

	return r.conn

}
func (r *Request) GetData() []byte {

	return r.data
}

func NewRequest(conn ziface.IConnection, data []byte) ziface.IRequest {

	r := &Request{
		conn: conn,
		data: data,
	}
	return r

}

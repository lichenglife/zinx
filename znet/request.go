package znet

import "lichenglife/zinx/ziface"

// 定义Request对象, 抽象连接以及请求数据
type Request struct {
	conn    ziface.IConnection
	message ziface.IMessage // 消息体
	// data []byte
}

func (r *Request) GetConn() ziface.IConnection {

	return r.conn

}
func (r *Request) GetData() []byte {

	return r.message.GetData()
}

func (r *Request) GetMsgID() uint32 {

	return r.message.GetMsgId()
}

func NewRequest(conn ziface.IConnection, message ziface.IMessage) ziface.IRequest {

	r := &Request{
		conn:    conn,
		message: message,
	}
	return r

}

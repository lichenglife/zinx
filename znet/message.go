package znet

import "lichenglife/zinx/ziface"

// 定义消息结构体

type Message struct {
	Data    []byte
	DataLen uint32
	Id      uint32
}

//  构造消息体

func NewMsgPackage(data []byte, datalen uint32, id uint32) ziface.IMessage {

	m := &Message{
		Data:    data,
		DataLen: datalen,
		Id:      id,
	}
	return m
}

// 获取消息长度
func (m *Message) GetDataLen() uint32 {
	return m.DataLen
}

// 获取消息ID

func (m *Message) GetMsgId() uint32 {

	return m.Id
}

// 获取数据
func (m *Message) GetData() []byte {

	return m.Data
}

// 设置消息ID

func (m *Message) SetMsgsId(id uint32) {

	m.Id = id
}

// 设置消息内容
func (m *Message) SetData(data []byte) {
	m.Data = data
}

// 设置消息长度

func (m *Message) SetMsgLen(len uint32) {
	m.DataLen = len
}

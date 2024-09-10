package ziface

/*
定义消息接口， 对请求中的消息体进行抽象封装
*/
type IMessage interface {

	// 获取消息长度
	GetDataLen() uint32
	// 获取消息ID

	GetMsgId() uint32
	// 获取数据
	GetData() []byte

	// 设置消息ID

	SetMsgsId(uint32)

	// 设置消息内容
	SetData([]byte)

	// 设置消息长度

	SetMsgLen(uint32)
}

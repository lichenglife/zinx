package ziface

type IRequest interface {

	//  获取连接

	GetConn() IConnection
	GetData() []byte
	GetMsgID() uint32
}

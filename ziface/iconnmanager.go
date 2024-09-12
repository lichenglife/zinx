package ziface

/*
	增加管理连接模块

server 端接收到客户端请求

将所有的connection 放在集合map[uint32]ziface.connection

增加连接管理

属性

方法
*/
type IConnManager interface {

	// 添加连接
	Add(conn IConnection)

	// 删除连接
	Remove(uint32)

	// 获取连接
	GetConn(uint32) (IConnection, error)

	// 获取当前连接池大小
	Len() int

	// 清空连接池
	ClearConn()
}

package ziface

//  定义Server的接口

type IServer interface {

	//  启动
	Start()

	//   停止
	Stop()
	//  运行
	Serve()

	// 添加路由函数
	AddRouter(msgID uint32, router IRouter)

	//  获取连接管理器
	GetConnMgr() IConnManager
}

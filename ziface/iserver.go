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

	AddRouter(router IRouter)
}

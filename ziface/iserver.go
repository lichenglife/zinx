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

	// Server添加注册Hook函数的方法， 在连接创建前、关闭前进行执行
	// 注册PreHook 函数

	SetOnConnStart(func(c IConnection))
	// 注册PostHook函数
	SetOnConnStop(func(c IConnection))
	// 执行preHook函数
	CallOnConnStart(c IConnection)
	// 执行PostHook 函数
	CallOnConnStop(c IConnection)
}

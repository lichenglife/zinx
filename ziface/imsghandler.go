package ziface

// 定义接口
// 抽象 消息以及处理消息的函数
type IMsgHandler interface {

	// 执行路由函数
	DoMsgHandler(request IRequest)

	// 注册路由
	AddRouter(msgID uint32, router IRouter)

	// 启动工作池 workPool中的 Request与handler

	StartWorkPool()

	// 将客户端请求添加到TaskQueue中
	SendMsgToTaskQueue(request IRequest)
}

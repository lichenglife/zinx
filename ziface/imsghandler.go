package ziface

// 定义接口
// 抽象 消息以及处理消息的函数
type IMsgHandler interface {

	// 执行路由函数
	DoMsgHandler(request IRequest)

	// 注册路由
	AddRouter(msgID uint32, router IRouter)
}

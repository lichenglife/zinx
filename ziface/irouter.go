package ziface

// 定义接口路由，
//
//	router的作用， 将connection与 处理业务函数的方法继续绑定
type IRouter interface {
	PreHandler(request IRequest)
	Handler(request IRequest)
	PostHandler(request IRequest)
}

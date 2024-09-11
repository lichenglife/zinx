package znet

import (
	"fmt"
	"lichenglife/zinx/utils"
	"lichenglife/zinx/ziface"
)

type MsgHandler struct {
	Apis map[uint32]ziface.IRouter

	// 业务工作的worke 工作池数量
	// worker 从队列中获取 客户端请求， 然后进行处理
	WorkPoolSize uint32

	//  定义消息队列,  业务请求的队列， 用于排队客户端的请求
	TaskQueue []chan ziface.IRequest
}

func NewMsgHandler() ziface.IMsgHandler {

	mh := &MsgHandler{
		Apis:         make(map[uint32]ziface.IRouter),
		WorkPoolSize: utils.GlobalObject.WorkPoolSize,
		TaskQueue:    make([]chan ziface.IRequest, utils.GlobalObject.WorkPoolSize),
	}
	return mh
}

// 注册路由
func (mh *MsgHandler) AddRouter(msgID uint32, router ziface.IRouter) {

	//  判断是否已经注册
	if _, ok := mh.Apis[msgID]; ok {
		panic(fmt.Sprintf("msgID  has registred %s", string(msgID)))
	}
	mh.Apis[msgID] = router

}

// 执行路由函数
func (mh *MsgHandler) DoMsgHandler(request ziface.IRequest) {

	handler, ok := mh.Apis[request.GetMsgID()]
	if !ok {
		fmt.Printf("msgId %d Handler is not registry", request.GetMsgID())
	}
	//  执行对应的处理
	handler.PreHandler(request)
	handler.Handler(request)
	handler.PostHandler(request)

}

// 启动Work ,  监听TaskQueue, 从TaskQueue中获取request， 然后执行DoMsgHandler,执行具体的handler
func (m *MsgHandler) StartWork(taskID int, taskQueue chan ziface.IRequest) {
	fmt.Printf("Start Task id is Start = %d \n", taskID)
	for {
		select {

		//	通过监听 chan 控制自动执行handler
		case request := <-taskQueue:
			//  执行具体的handler
			m.DoMsgHandler(request)
		}
	}
}

func (mh *MsgHandler) StartWorkPool() {

	//  启动与配置数量相同的Work

	for i := 0; i < int(utils.GlobalObject.WorkPoolSize); i++ {

		mh.TaskQueue[i] = make(chan ziface.IRequest, utils.GlobalObject.MaxWorkerTaskLen)
		go mh.StartWork(i, mh.TaskQueue[i])
	}

}

// 将客户端请求添加到TaskQueue中
func (mh *MsgHandler) SendMsgToTaskQueue(request ziface.IRequest) {

	//  通过取模计算workID 应该存在哪一个TaskQueue中
	workID := request.GetConn().GetConnID() % mh.WorkPoolSize
	// 切片类型
	mh.TaskQueue[workID] <- request

}

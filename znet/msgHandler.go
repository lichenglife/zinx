package znet

import (
	"fmt"
	"lichenglife/zinx/ziface"
)

type MsgHandler struct {
	Apis map[uint32]ziface.IRouter
}

func NewMsgHandler() ziface.IMsgHandler {

	mh := &MsgHandler{
		Apis: make(map[uint32]ziface.IRouter),
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

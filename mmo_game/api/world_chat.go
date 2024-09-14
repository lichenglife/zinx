package api

import (
	"fmt"
	"lichenglife/zinx/mmo_game/core"
	"lichenglife/zinx/mmo_game/pb"
	"lichenglife/zinx/ziface"
	"lichenglife/zinx/znet"

	"google.golang.org/protobuf/proto"
)

// 玩家聊天业务函数

//  注册路由函数实现业务

type WorldChatApi struct {
	znet.BaseRouter
}

func (wcr *WorldChatApi) Handler(request ziface.IRequest) {

	msg := &pb.Talk{}
	// 反序列化用户数据
	err := proto.Unmarshal(request.GetData(), msg)
	if err != nil {
		fmt.Printf("unmarshal data error %v", err)
		return
	}
	pid, err := request.GetConn().GetProperty("pid")
	if err != nil {
		fmt.Printf("Get pid from request connection error %v", err)
		request.GetConn().Stop()
		return
	}
	fmt.Println("current player pid is ", pid)

	player := core.WorldMgrObj.GetPlayerByPid(pid.(int32))

	player.Talk(msg.Content)
}

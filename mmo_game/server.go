package main

import (
	"fmt"
	"lichenglife/zinx/mmo_game/api"
	"lichenglife/zinx/mmo_game/core"
	"lichenglife/zinx/ziface"
	"lichenglife/zinx/znet"
)

// 用户上线通知
func OnConnectionAction(conn ziface.IConnection) {

	player := core.NewPlayer(conn)

	player.SyncPid()
	// 同步当前玩家的初始化坐标信息给客户端，走MsgID:200消息
	player.BroadCastStartPosition()

	//  将当前玩家添加到WorldManager 中
	core.WorldMgrObj.AddPlayer(player)

	//=================将该连接绑定属性Pid===============
	conn.SetProperty("pid", player.Pid)

	fmt.Println("=====> Player pidId = ", player.Pid, " arrived ====")

	//==============同步周边玩家上线信息，与现实周边玩家信息========
	player.SyncSurrounding()

}

func OnConnectionLost(conn ziface.IConnection) {
	//获取当前连接的Pid属性
	pid, _ := conn.GetProperty("pid")

	//根据pid获取对应的玩家对象
	player := core.WorldMgrObj.GetPlayerByPid(pid.(int32))

	//触发玩家下线业务
	if pid != nil {
		player.LostConnection()
	}

	fmt.Println("====> Player ", pid, " left =====")

}

// 客户端入口
func main() {

	// 1、创建Server
	s := znet.NewServer()

	s.SetOnConnStart(OnConnectionAction)
	s.SetOnConnStop(OnConnectionLost)

	//注册路由
	s.AddRouter(2, &api.WorldChatApi{}) //聊天
	s.AddRouter(3, &api.MoveApi{})      //移动

	// 2、启动Server
	s.Serve()
}

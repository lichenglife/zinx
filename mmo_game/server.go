package main

import (
	"fmt"
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

	fmt.Println("=====> Player pidId = ", player.Pid, " arrived ====")

}

// 客户端入口
func main() {

	// 1、创建Server
	s := znet.NewServer()

	s.SetOnConnStart(OnConnectionAction)

	// 2、启动Server
	s.Serve()
}

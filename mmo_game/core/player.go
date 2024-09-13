package core

import (
	"fmt"
	"lichenglife/zinx/mmo_game/pb"
	"lichenglife/zinx/ziface"
	"math/rand"
	"sync"

	"google.golang.org/protobuf/proto"
)

type Player struct {
	Pid  int32              // 玩家ID
	Conn ziface.IConnection // 当前连接
	X    float32            // x坐标轴位置
	Y    float32            // Y坐标轴位置
	Z    float32            // 高度
	V    float32            // 角度
}

/*
Player ID 生成器
*/
var PidGen int32 = 1  //用来生成玩家ID的计数器
var IdLock sync.Mutex //保护PidGen的互斥机制

// 创建一个玩家对象
func NewPlayer(conn ziface.IConnection) *Player {
	//生成一个PID
	IdLock.Lock()
	id := PidGen
	PidGen++
	IdLock.Unlock()

	p := &Player{
		Pid:  id,
		Conn: conn,
		X:    float32(160 + rand.Intn(10)), //随机在160坐标点 基于X轴偏移若干坐标
		Y:    0,                            //高度为0
		Z:    float32(134 + rand.Intn(17)), //随机在134坐标点 基于Y轴偏移若干坐标
		V:    0,                            //角度为0，尚未实现
	}

	return p
}

// 发送消息   protobuff 格式的消息内容

func (p *Player) SendMsg(msgId uint32, data proto.Message) {
	fmt.Printf("before Marshal data = %+v\n", data)
	// 序列化
	msg, err := proto.Marshal(data)
	if err != nil {
		fmt.Printf("marshal msg  error %v \n", err)
		return
	}
	fmt.Printf("after Marshal data = %+v\n", msg)

	if p.Conn == nil {
		fmt.Printf("player %d  conn is nill \n", p.Pid)
		return
	}

	if err := p.Conn.SendMsg(msgId, msg); err != nil {
		fmt.Printf("player %d  send msg  error %v \n", p.Pid, err)
		return
	}

}

// 告知客户端pid,同步已经生成的玩家ID给客户端
func (p *Player) SyncPid() {
	// 组建MsgId0 proto数据
	// protobuff  消息结构体
	data := &pb.SyncPid{
		Pid: p.Pid,
	}

	//发送数据给客户端
	p.SendMsg(1, data)

}

func (p *Player) BroadCastStartPosition() {

	msg := &pb.BroadCast{
		Pid: p.Pid,
		Tp:  2, //TP2 代表广播坐标
		Data: &pb.BroadCast_P{
			P: &pb.Position{
				X: p.X,
				Y: p.Y,
				Z: p.Z,
				V: p.V,
			},
		},
	}

	p.SendMsg(200, msg)

}

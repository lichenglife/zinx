package main

import (
	"fmt"
	"io"
	"lichenglife/zinx/znet"
	"net"
	"time"
)

func main() {

	time.Sleep(2 * time.Second)

	//  客户端代码
	fmt.Println("Client Start ....")
	conn, err := net.Dial("tcp", "127.0.0.1:8999")
	if err != nil {
		fmt.Println("Client Start err ", err)
		return
	}

	for {

		// 封包 发送数据
		dp := znet.NewDataPack()
		msg, _ := dp.Pack(znet.NewMsgPackage([]byte("=====  i am robot  number  two  ====="), 2))
		_, err := conn.Write(msg)
		if err != nil {
			fmt.Println("write error err ", err)
			return
		}

		// 解包接受数据

		//先读出流中的head部分
		headData := make([]byte, dp.GetHeadLen())
		_, err = io.ReadFull(conn, headData) //ReadFull 会把msg填充满为止
		if err != nil {
			fmt.Println("read head error")
			break
		}

		//将headData字节流 拆包到msg中
		msgHead, err := dp.UpPack(headData)
		if err != nil {
			fmt.Println("server unpack err:", err)
			return
		}

		if msgHead.GetDataLen() > 0 {
			//msg 是有data数据的，需要再次读取data数据
			msg := msgHead.(*znet.Message)
			msg.Data = make([]byte, msg.GetDataLen())

			//根据dataLen从io中读取字节流
			_, err := io.ReadFull(conn, msg.Data)
			if err != nil {
				fmt.Println("server unpack data err:", err)
				return
			}

			fmt.Println("==> Recv Msg: ID=", msg.Id, ", len=", msg.DataLen, ", data=", string(msg.Data))
		}
	}
	time.Sleep(1 * time.Second)

}

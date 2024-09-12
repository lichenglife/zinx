package main

import (
	"fmt"
	"io"
	"lichenglife/zinx/znet"
	"net"
)

func main() {

	//  创建Tcp socket
	listener, err := net.Listen("tcp", "127.0.0.1:8990")
	if err != nil {
		fmt.Println("test  error", err)
		return
	}
	for {

		//  测试解包

		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Accept error", err)
			return
		}

		go func(conn net.Conn) {
			// 创建封包解包对象
			dp := &znet.DataPack{}

			for {
				// 循环读取数据
				headata := make([]byte, dp.GetHeadLen())
				cnt, err := io.ReadFull(conn, headata)
				if err != nil {
					return
				}
				fmt.Println(cnt)

				msgHead, err := dp.UpPack(headata)
				if err != nil {
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

		}(conn)

	}

}

//  业务抽象

//  代码规范

//  行为思路

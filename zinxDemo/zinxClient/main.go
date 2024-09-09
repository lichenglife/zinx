package main

import (
	"fmt"
	"net"
	"time"
)

func main() {

	for i := 0; i < 10; i++ {

		time.Sleep(2 * time.Second)

		//  客户端代码
		fmt.Println("Client Start ....")
		conn, err := net.Dial("tcp", "127.0.0.1:8999")
		if err != nil {
			fmt.Println("Client Start err ", err)
			return
		}

		for {
			//TODO
			_, err := conn.Write([]byte("Hello Zinx Server , I am Zinx Client V0.1"))
			if err != nil {
				fmt.Println("writer to server errror", err)
				return
			}

			// 定义Request Response
			data := make([]byte, 512)

			cnt, err := conn.Read(data)
			if err != nil {
				fmt.Println("read  from server error", err)
				return
			}

			fmt.Printf("read from server message %s, and cnt %d \n", string(data), cnt)

		}
	}
}

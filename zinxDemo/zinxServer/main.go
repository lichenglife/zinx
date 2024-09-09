package main

import (
	"fmt"
	"lichenglife/zinx/znet"
)

func main() {

	fmt.Println("Start  Zinx Server v0.1")

	s := znet.NewServer("zinServer", "127.0.0.1", 8999)

	s.Serve()

}

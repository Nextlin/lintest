package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	//创建socke后，定义使用完成后的关闭
	defer conn.Close()
	if err != nil {
		fmt.Println("Dial err:", err)
		return
	}
	//主动发送数据给服务器
	conn.Write([]byte("this is the client massage"))
	//定义一个切片接收服务器返回的内容
	buf := make([]byte, 4096)
	n, err1 := conn.Read(buf)
	if err1 != nil {
		fmt.Println("conn.Read err:", err)
		return
	}
	fmt.Println("服务器端返回的内容:", string(buf[:n]))
}

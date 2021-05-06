package main

import (
	"fmt"
	"net"
)

func main() {
	//指定监听服务器端的协议，ip端口
	Linsner, err := net.Listen("tcp", "127.0.0.1:8000")
	defer Linsner.Close()
	if err != nil {
		fmt.Println("net.Listen有错误，启动失败", err)
		return
	}
	fmt.Print("等待建立连接")
	//阻塞客户端，等待客服端链接
	Conne, err1 := Linsner.Accept()
	defer Conne.Close()
	if err1 != nil {
		fmt.Print("Linsner.Accept() eror:", err1)
		return
	}
	fmt.Println("服务端与客服端建立了链接")
	buf := make([]byte, 4096)
	n, err2 := Conne.Read(buf)
	if err2 != nil {
		fmt.Print("Conne.Read() eror:", err1)
		return
	}
	fmt.Println("客服端的输入是：", string(buf[:n]))
}

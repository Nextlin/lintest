package main

import (
	"fmt"
	"net"
	"strings"
)

func HandelCon(con net.Conn) {
	addr := con.RemoteAddr()
	fmt.Println("客户端的ip是：", addr)
	//在本次链接中循环的读取客服端发送的内容。1.定义一个切片接收单词的请求，2.循环的从切片读取内容
	buf := make([]byte, 4096)
	for {
		n, err2 := con.Read(buf)
		if err2 != nil && n != 0 {
			fmt.Println("con.Read error", err2)
			return
		} else if n == 0 {
			fmt.Println("客服端自行关闭")
			return
		}
		fmt.Println("服务器读到的客户端发送的数据是:", string(buf[:n]))
		//将读到的输入做一下处理，比如小写转大写strings.ToUpper(string(buf[:n])),Write方法是服务端回复客服端的方法
		con.Write([]byte(strings.ToUpper(string(buf[:n]))))
	}
	defer con.Close()
}

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	fmt.Println("开始了")
	if err != nil {
		fmt.Println("Listen ERROR:", err)
		return
	}
	defer listener.Close() //在完成后，随后将监听关闭

	for {
		con, err1 := listener.Accept()
		fmt.Println("链接建立了")
		if err1 != nil {
			fmt.Println("listener.Accept error", err1)
			return
		}
		go HandelCon(con)
	}

}

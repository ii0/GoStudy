package main

import (
	"fmt"
	"net"
	"log"
	"os"
)

func main() {

	//建立socket，监听端口  第一步:绑定端口
	netListen, err := net.Listen("tcp", "localhost:1024")
	CheckError(err)
	defer netListen.Close()

	Log("Waiting for clients")
	for {
		conn, err := netListen.Accept()  //第二步:获取连接
		if err != nil {
			continue  //出错退出当前一次循环
		}

		Log(conn.RemoteAddr().String(), " tcp connect success")
		handleConnection(conn)  //正常连接就处理
	}
}
//处理连接
func handleConnection(conn net.Conn) {

	buffer := make([]byte, 2048)

	for {  //无限循环

		n, err := conn.Read(buffer) //第三步:读取从该端口传来的内容

		if err != nil {
			Log(conn.RemoteAddr().String(), " connection error: ", err)
			return //出错后返回
		}

		Log(conn.RemoteAddr().String(), "receive data string:\n", string(buffer[:n]))

	}
}

//log输出
func Log(v ...interface{}) {
	log.Println(v...)
}

//处理error
func CheckError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
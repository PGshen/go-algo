package usepackage

import (
	"fmt"
	"net"
	"os"
)

func tcpRecv() {
	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("net.Listen err:", err)
		return
	}
	defer listener.Close()

	fmt.Println("waiting to recv...")

	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("listener.Accept err: ", err)
		return
	}
	defer conn.Close()

	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("conn.Read err: ", err)
		return
	}
	fileName := string(buf[:n])
	conn.Write([]byte("ok"))
	recvFile(conn, fileName)
}

// 接收文件
func recvFile(conn net.Conn, fileName string) {
	f, err := os.Create(fileName)
	if err != nil {
		fmt.Println("os.Create err: ", err)
		return
	}
	defer f.Close()

	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("conn.Read err: ", err)
			return
		}
		if n == 0 {
			fmt.Println("receive file finish!")
			return
		}
		f.Write(buf[:n])
	}
}

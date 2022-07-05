package usepackage

import (
	"fmt"
	"io"
	"net"
	"os"
)

func tcpSend() {
	list := os.Args
	if len(list) != 2 {
		fmt.Println("format: go run xx.go /absolute/path")
		return
	}
	filePath := list[1]
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		fmt.Println("os.Stat err: ", err)
		return
	}
	fileName := fileInfo.Name()

	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("net.Dial err: ", err)
		return
	}
	defer conn.Close()

	// 发送文件名
	_, err = conn.Write([]byte(fileName))
	if err != nil {
		fmt.Println("conn.Write err: ", err)
		return
	}
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("conn.Read err: ", err)
		return
	}
	// 发送文件
	if "ok" == string(buf[:n]) {
		sendFile(conn, filePath)
	}
}

func sendFile(conn net.Conn, filePath string) {
	// 打开文件
	f, err := os.Open(filePath)
	if err != nil {
		fmt.Println("os.Open err: ", err)
		return
	}
	defer f.Close()

	buf := make([]byte, 1024)
	for {
		n, err := f.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Println("send file finish!")
			} else {
				fmt.Println("os.Read err: ", err)
			}
		}
		_, err = conn.Write(buf[:n])
		if err != nil {
			fmt.Println("conn.Write err: ", err)
			return
		}
	}
}

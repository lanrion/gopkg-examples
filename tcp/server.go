package main

import (
	"fmt"
	"net"
	"strings"
	"time"
)

var conns = make(map[string]net.Conn)

func main() {
	fmt.Println("Starting the server ...")
	// 创建 listener
	listener, err := net.Listen("tcp", "192.168.0.182:50000")
	if err != nil {
		fmt.Println("Error listening", err.Error())
		return //终止程序
	}
	// 监听并接受来自客户端的连接
	for {
		conn, err := listener.Accept()

		if err != nil {
			fmt.Println("Error accepting", err.Error())
			return // 终止程序
		}
		go doServerStuff(conn)
		go doNotice()
	}
}

func doNotice() {
	for {
		time.Sleep(2 * time.Second)
		for client, conn := range conns {
			fmt.Println("Notice client: ", client)
			_, er := conn.Write([]byte("12312"))
			if er != nil {
				conn.Close()
				fmt.Println("Error reading", er.Error())
			}
		}
	}
}

func doServerStuff(conn net.Conn) {
	for {
		buf := make([]byte, 512)
		len, err := conn.Read(buf)

		if err != nil {
			fmt.Println("Error reading", err.Error())
			return //终止程序
		}

		msg := string(buf[:len])

		clientName := strings.TrimSpace(strings.Split(msg, "says")[0])

		if conns[clientName] == nil {
			conns[clientName] = conn
		}

		fmt.Printf("Received data: %v", msg)
	}
}

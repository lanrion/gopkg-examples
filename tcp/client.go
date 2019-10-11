package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	//打开连接:
	conn, err := net.Dial("tcp", "192.168.0.182:50000")

	if err != nil {
		//由于目标计算机积极拒绝而无法创建连接
		fmt.Println("Error dialing", err.Error())
		return // 终止程序
	}

	go doRead(conn)

	for {
		doWrite(conn)
	}

}

func doWrite(conn net.Conn) {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("First, what is your name?")
	clientName, _ := inputReader.ReadString('\n')

	trimmedClient := strings.Trim(clientName, "\r\n")
	fmt.Println("What to send to the server? Type Q to quit.")
	input, _ := inputReader.ReadString('\n')
	trimmedInput := strings.Trim(input, "\r\n")

	if trimmedInput == "Q" {
		return
	}
	_, err := conn.Write([]byte(trimmedClient + " says: " + trimmedInput))
	if err != nil {
		fmt.Println("Error reading", err.Error())
		return //终止程序
	}
}

func doRead(conn net.Conn) {
	for {
		buf := make([]byte, 512)
		len, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error reading", err.Error())
			return //终止程序
		}
		fmt.Printf("Received Response: %v", string(buf[:len]))
		fmt.Println()
	}

}

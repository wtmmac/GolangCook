package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println(Show("\n输入的参数不正确,请重新输入.\r\n可选参数server或client"))
		os.Exit(0)
	}
	param := strings.ToLower(args[1])
	if param == "server" {
		Server()
	} else if param == "client" {
		Client()
	} else {
		fmt.Println(Show("输入的参数不正确,请重新输入.\r\n可选参数server或client"))
		os.Exit(0)
	}
}

// 编辑转换（Todo）
func Show(s string) string {
	return s
}

func Server() {
	exit := make(chan bool)
	ip := net.ParseIP("127.0.0.1")
	addr := net.TCPAddr{ip, 9872, ""}
	go func() {
		listen, err := net.ListenTCP("tcp", &addr)
		if err != nil {
			fmt.Println(Show("初始化失败"), Show(err.Error()))
			exit <- true
			return
		}
		fmt.Println(Show("Listening..."))
		for {
			client, err := listen.AcceptTCP()
			if err != nil {
				fmt.Println(err.Error())
				continue
			}
			fmt.Println(Show("Client Connected"), Show(client.RemoteAddr().String()))
			data := make([]byte, 1024)
			c, err := client.Read(data)
			if err != nil {
				fmt.Println(Show(err.Error()))
			}
			fmt.Println(string(data[0:c]))
			fmt.Println("收到:", c, " bytes\n")
			client.Write([]byte("你好客户端!\r\n"))
			client.Close()
		}
	}()
	<-exit
	fmt.Println(Show("服务端关闭!"))
}

func Client() {
	client, err := net.Dial("tcp", "127.0.0.1:9872")
	if err != nil {
		fmt.Println(Show("服务端连接失败"), Show(err.Error()))
		return
	}
	defer client.Close()
	client.Write([]byte("你好,服务端!"))
	buf := make([]byte, 1024)
	c, err := client.Read(buf)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(Show(string(buf[0:c])))
}

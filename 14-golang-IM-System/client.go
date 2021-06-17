package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"strconv"
)

type Client struct {
	ServerIp 	string
	ServerPort  int
	Name 		string
	conn 		net.Conn
	flag 		int // 当前client的模式
}

func NewClient(serverIp string, serverPort int) *Client {
	// 创建客户端对象
	client := &Client {
		ServerIp: 	serverIp,
		ServerPort: serverPort,
		flag: 		999,
	}

	// 连接server
	// 客户端连接服务器接口，返回通信套接字
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", serverIp, serverPort))
	if err != nil {
		fmt.Println("net.Dial error:", err)
		return nil
	}

	client.conn = conn

	// 返回对象
	return client
}

func (client *Client) menu() bool {
	var input string
	fmt.Println("1.群聊模式")
	fmt.Println("2.私聊模式")
	fmt.Println("3.更新用户名")
	fmt.Println("0.退出")

	// 从终端读取数据
	fmt.Scanln(&input)

	flag, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println(">>>>>请输入合法范围内的数字(0~3)")
		return false
	}

	if flag >= 0 && flag <= 3 {
		client.flag = flag
		return true
	} else {
		fmt.Println(">>>>>请输入合法范围内的数字(0~3)")
		return false
	}
}

// 查询在线用户
func (client *Client) SelectUsers() {
	sendMsg := "who\n"
	_, err := client.conn.Write([]byte(sendMsg))
	if err != nil {
		fmt.Println("conn.Write error:", err)
		return
	}
}

func (client *Client) PrivateChat() {
	var remoteName string
	var chatMsg string

	client.SelectUsers()

	fmt.Println(">>>>>请输入聊天对象的用户名，exit退出")
	fmt.Scanln(&remoteName)

	for remoteName != "exit" {
		fmt.Println(">>>>>请输入消息内容，exit退出")
		fmt.Scanln(&chatMsg)

		for chatMsg != "exit" {
			// 消息不为空则发送
			if len(chatMsg) != 0 {
				sendMsg := "to|" + remoteName + "|" + chatMsg + "\n"
				_, err := client.conn.Write([]byte(sendMsg))
				if err != nil {
					fmt.Println("conn.Write error:", err)
					break
				}
			}

			//chatMsg = "" // 消息重置为空，有点多余
			fmt.Println(">>>>>请输入聊天内容，exit退出") // 继续提醒输入
			fmt.Scanln(&chatMsg) // 从控制台阻塞读
		}

		client.SelectUsers()
		fmt.Println(">>>>>请输入聊天对象的用户名，exit退出")
		fmt.Scanln(&remoteName)
	}
}

func (client *Client) PublicChat() {
	// 提示用户输入消息
	var chatMsg string
	fmt.Println(">>>>>请输入聊天内容，exit退出")
	fmt.Scanln(&chatMsg)

	// 循环提示输入消息，并监听控制台输入数据
	for chatMsg != "exit" {
		// 发给服务器

		// 消息不为空则发送
		if len(chatMsg) != 0 {
			sendMsg := chatMsg + "\n"
			_, err := client.conn.Write([]byte(sendMsg))
			if err != nil {
				fmt.Println("conn.Write error:", err)
				break
			}
		}

		// chatMsg = "" // 消息重置为空，有点多余
		fmt.Println(">>>>>请输入聊天内容，exit退出") // 继续提醒输入
		fmt.Scanln(&chatMsg) // 从控制台阻塞读
	}
}

func (client *Client) UpdateName() bool {
	fmt.Println(">>>>>请输入用户名：")
	fmt.Scanln(&client.Name) // User内数据修改

	sendMsg := "rename|" + client.Name + "\n"
	_, err := client.conn.Write([]byte(sendMsg)) // 服务器内数据修改
	if err != nil {
		fmt.Println("conn.Write error:", err)
		return false
	}
	return true
}

// 处理server回应的消息，直接显示在屏幕上
func(client *Client) DealResponse() {
	// 一旦client.conn有数据，就直接拷贝到stdout标准输出，永久监听阻塞
	io.Copy(os.Stdout, client.conn) // 阻塞写，相当于下面的for循环
	//for {
	//	buf := make()
	//	client.conn.Read(buf)
	//	fmt.Println(buf)
	//}
}

func (client *Client) Run() {
	// for关键字代替while循环
	for client.flag != 0 {
		// 输入错误，不执行任何操作，反复输出提示信息
		for client.menu() != true {}

		// 根据不同模式处理不同的业务
		switch client.flag {
		case 1:
			// 群聊模式
			client.PublicChat()
			break
		case 2:
			// 私聊模式
			client.PrivateChat()
		case 3:
			// 更新用户名
			client.UpdateName()
			break
		}
	}
}

var serverIp string
var serverPort int

// 在main函数之前执行
// 内部解析命令行，解析出的值传给serverIp和serverPort
func init() {
	// ./client -ip 127.0.0.1 8888
	// -h 可以显示提示信息
	flag.StringVar(&serverIp, "ip", "127.0.0.1", "服务器IP地址，默认是127.0.0.1")
	flag.IntVar(&serverPort, "port", 8888, "服务器端口，默认是8888")
}

func main() {
	// 命令行解析
	flag.Parse()

	client := NewClient(serverIp, serverPort)
	if client == nil {
		fmt.Println(">>>>>连接服务器失败...")
		return
	}

	// 主go程阻塞显示菜单，开一个子go程显示服务器的回执消息
	// 这里只是承载，并没有执行
	go client.DealResponse()

	fmt.Println(">>>>>连接服务器成功...")

	// 启动客户端的业务
	client.Run()
}
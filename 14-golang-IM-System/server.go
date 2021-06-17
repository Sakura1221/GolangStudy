package main

import (
	"fmt"
	"io"
	"net"
	"sync"
	"time"
)

type Server struct {
	Ip string
	Port int

	// 在线用户的列表
	OnlineMap map[string]*User
	mapLock sync.RWMutex // 读写锁

	// 消息广播的channel
	Message chan string
}

// 创建一个server的接口
func NewServer(ip string, port int) *Server {
	// 初始化
	server := &Server {
		Ip: ip,
		Port: port,
		OnlineMap: make(map[string]*User),
		Message: make(chan string),
	}
	return server
}

// 监听Message广播消息channel的goroutine，一旦有消息就发送给全部的在线user
func (this *Server) ListenMessager() {
	for {
		msg := <-this.Message

		// 将msg发送给全部的在线User
		this.mapLock.Lock()
		for _, client := range this.OnlineMap {
			client.C <- msg
		}
		this.mapLock.Unlock()
	}
}

// 广播消息的方法
func (this *Server) BroadCast(user *User, msg string) {
	sendMsg := "[" + user.Addr + "]" + user.Name + ":" + msg

	// 向广播通道发送消息
	this.Message <- sendMsg
}

func (this *Server) Handler(conn net.Conn) {
	user := NewUser(conn, this)

	// 用户上线业务，封装在用户类中
	user.Online()

	// 监听用户是否活跃的channel
	isLive := make(chan bool)

	// 接收客户端发送到消息
	go func() {
		buf := make([]byte, 4096)
		for {
			n, err := conn.Read(buf)

			// 异常
			if err != nil && err != io.EOF {
				fmt.Println("conn read error:", err)
				return
			}

			// 客户端正常关闭
			if n == 0 {
				// 用户下线业务，封装在用户类中
				user.Offline()
				return
			}

			// 提取用户的消息（取出'\n'）
			msg := string(buf[:n-1])

			// 将得到的消息进行处理
			user.DoMessage(msg)

			// 用户发送任意消息，代表当前用户是一个活跃的
			isLive <- true
		}
	}()

	for {
		// 当前handler阻塞
		select {
			// 如果用户不活跃，且未超时，那么所有channel都未就绪
			// select会阻塞，直到用户活跃，或者超时
			case <- isLive:
				// 当前用户是活跃的，应该重置定时器
				// 不做任何事情，为了激活select，更新下面的定时器

			// go语言中的定时器，本质是一个channel
			// 每执行一次select，定时器都会重置
			case <- time.After(time.Second * 600):
				// 已经超时
				// 将当前的User强制关闭
				user.SendMsg("你长时间不操作，被踢了")

				// 用户下线
				user.Offline()

				// 销毁相关资源
				close(user.C)

				// 关闭连接
				conn.Close()

				// 退出当前Handler（Go程）
				return // runtime.Goexit()
		}
	}
}

// 启动服务器的接口
func (this *Server) Start() {
	// socket listen
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", this.Ip, this.Port))
	if err != nil {
		fmt.Println("net listen error:", err)
		return
	}
	// close listen socket
	defer listener.Close() // 提前写避免忘记关闭

	// 启动监听Message的goroutine（监听写）
	go this.ListenMessager()

	// 服务器循环监听并处理业务
	for {
		// accept
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener accept error:", err)
			continue
		}

		// do handler
		// 用一个go程处理业务
		go this.Handler(conn)
	}
}

package main

import (
	"net"
	"strings"
)

type User struct {
	Name string
	Addr string
	C chan string // 管道作为发送数据的缓冲区，阻塞读数据发送
	conn net.Conn // 用来通信的连接

	server *Server // 当前用户属于哪个server，负责执行业务
}

// 创建一个用户的API
func NewUser(conn net.Conn, server *Server) *User {
	userAddr := conn.RemoteAddr().String()

	user := &User {
		Name: userAddr,
		Addr: userAddr,
		C:    make(chan string),
		conn: conn,

		server : server,
	}

	// 启动监听当前user channel消息的goroutine（监听读）
	go user.ListenMessage()

	return user
}

// 用户的上线业务
func (this *User) Online() {
	// 用户上线，将用户加入到onlineMap中
	this.server.mapLock.Lock()
	this.server.OnlineMap[this.Name] = this
	this.server.mapLock.Unlock()

	// 广播当前用户上线消息
	this.server.BroadCast(this, "已上线")
}

// 用户的下线业务
func (this *User) Offline() {
	// 用户下线，将用户从onlineMap中删除
	this.server.mapLock.Lock()
	delete(this.server.OnlineMap, this.Name)
	this.server.mapLock.Unlock()

	// 广播当前用户下线消息
	this.server.BroadCast(this, "已下线")
}

// 给当前User对应的客户端发送消息
func (this *User) SendMsg(msg string) {
	this.conn.Write([]byte(msg))
}

// 处理用户消息的业务
func (this *User) DoMessage(msg string) {
	// 查询当前在线用户都有哪些
	if msg == "who" {
		this.server.mapLock.Lock()
		for _, user := range this.server.OnlineMap {
			onlineMsg := "[" + user.Addr + "]" + user.Name + ":" + "在线...\n"
			this.SendMsg(onlineMsg)
		}
		this.server.mapLock.Unlock()
	} else if len(msg) > 7  && msg[:7] == "rename|" {
		// 消息格式： rename|张三
		newName := strings.Split(msg, "|")[1] // 字符串分割

		// 判断name是否存在
		_, ok := this.server.OnlineMap[newName]
		if ok {
			this.SendMsg("当前用户名被使用\n")
		} else {
			this.server.mapLock.Lock()
			delete(this.server.OnlineMap, this.Name)
			this.server.OnlineMap[newName] = this
			this.server.mapLock.Unlock()

			this.Name = newName
			this.SendMsg("您已经更新用户名：" + this.Name + "\n")
		}
	} else if len(msg) > 4 && msg[:3] == "to|" {
		// 消息格式： to|张三|消息内容

		// 1 获取对方的用户名，和要发送的消息
		s := strings.Split(msg, "|")
		remoteName := s[1]
		content := s[2]
		if remoteName == "" {
			this.SendMsg("消息格式不正确，请使用\"to|张三|你好啊\"格式\n")
			return // 错误情况，记得return
		}
		if content == "" {
			this.SendMsg("发送消息不能为空")
			return
		}

		// 3 根据用户名得到对方User对象，将消息发送过去
		remoteUser, ok := this.server.OnlineMap[remoteName]
		if !ok {
			this.SendMsg("该用户名不存在\n")
			return
		}
		remoteUser.SendMsg(this.Name + "对您说：" + content + "\n")
	} else {
		// 将得到的用户消息进行广播
		this.server.BroadCast(this, msg)
	}
}

// 监听当前User channel的方法，一旦有消息就直接发送给对端客户端
func (this *User) ListenMessage() {
	for {
		msg := <-this.C

		this.conn.Write([]byte(msg + "\n"))
	}
}

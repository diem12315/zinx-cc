package znet

import (
	"errors"
	"fmt"
	"net"

	"example.com/m/ziface"
)

type Server struct {
	//服务器的名称
	Name string
	//服务器绑定的IP
	IPVersion string
	//服务器监听的IP
	IP string
	// 服务器监听的端口
	Port int
}

// 定义客户端链接所绑定的handle api(目前是固定的，后期由用户修改)
func CallBackClient(conn *net.TCPConn, data []byte, cnt int) error {
	// 回显业务
	fmt.Println("[Conn Handle] CallbackToClient ....")
	if _, err := conn.Write(data[:cnt]); err != nil {
		fmt.Println("write back buf err", err)
		return errors.New("CallBack To Client error")
	}

	return nil
}

// 启动服务器
func (s *Server) Start() {
	fmt.Printf("[start] Server Linstnner at %s:%d, is starting\n", s.IP, s.Port)
	go func() {
		// 1. 获取一个tcp的addr
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("resolve tcp addt error :", err)
			return
		}

		// 2. 监听服务器的地址
		listen, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Println("listen ", s.IPVersion, "err", err)
			return
		}

		fmt.Println("start Zinx server succ,", s.Name, "succ,Listennning..")

		// 变量和赋值可以在同一行
		var cid uint32 = 0

		// 3. 阻塞的等待客户端连接，处理客户端链接业务(读写)
		for {
			// 如果有客户端链接过来，阻塞会返回
			conn, err := listen.AcceptTCP()
			if err != nil {
				fmt.Println("Accept err", err)
				continue
			}
			// 将处理新链接的业务方法和conn进行绑定得到我们的链接模块
			dealConn := NewConnection(conn, cid, CallBackClient)
			cid++

			// 启动当前链接业务处理
			go dealConn.Start()
		}
	}()
}

// 停止服务器
func (S *Server) Stop() {
	// TODO 将一些服务器的资源，状态或者一些已经开辟的链接信息 进行停止或是回收

}

// 运行服务器
func (s *Server) Server() {
	// 启动server的服务功能
	s.Start()

	// TODO 做一些启动服务器之后的额外业务

	// 阻塞状态
	select {}
}

/*
	初始Server模块方法
*/
func NewServer(name string) ziface.IServer {
	s := &Server{
		Name:      name,
		IPVersion: "tcp4",
		IP:        "0.0.0.0",
		Port:      8999,
	}

	return s
}

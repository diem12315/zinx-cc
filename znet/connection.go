package znet

import (
	"net"

	"example.com/m/ziface"
)

/*
	链接模块
*/
type Connection struct {
	// 当前链接的socket TCP套接字
	Conn *net.TCPConn
	// 当前链接的ID
	ConnID uint32
	// 当前的链接状态
	isClosed bool
	// 当前链接所绑定的处理业务方法API
	handleAPI ziface.HandleFunc
	// 告知当前链接已经退出的/停止 channel
	ExitChan chan bool
}

// 初始化链接模块的方法

func NewConnection(conn *net.TCPConn, connID uint32, callback_api ziface.HandleFunc) *Connection {
	c := &Connection{
		Conn:      conn,
		ConnID:    connID,
		handleAPI: callback_api,
		isClosed:  false,
		ExitChan:  make(chan bool, 1),
	}

	return c
}

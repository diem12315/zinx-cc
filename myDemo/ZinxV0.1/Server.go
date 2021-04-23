package main

import "example.com/m/znet"

/*
	基于Zinx框架开发的 服务器端应用程序
*/

func main() {

	//1.创建一个server句柄，使用Zinx的api
	s := znet.NewServer("[zinx v0.1]")
	//2.启动server
	s.Server()

}

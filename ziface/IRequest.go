package ziface

/*
	IReqeust接口
	实际上是把客户端请求的链接学习，和请求的数据包装到一个Request
*/

type IRequeust interface {
	//得到当前链接
	GetConnection() IConnection
	// 得到请求的消息数据
	GetData() []byte
}

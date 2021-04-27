package znet

import "example.com/m/ziface"

// 实现router时，先嵌入这个baseRouter基类，然后根据需要对这个基类方法进行重写。
type BaseRouter struct {
}

// 在处理conn业务之前的钩子方法Hook
func (br *BaseRouter) PreHandle(request ziface.IRequeust) {

}

// 在处理conn业务的主方法Hook
func (br *BaseRouter) Handle(request ziface.IRequeust) {

}

// 在处理conn业务之后的钩子方法hook
func (br *BaseRouter) PostHandle(request ziface.IRequeust) {

}

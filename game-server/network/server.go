package network

import (
	"fmt"
	"net"
)

type Server struct {
	listener        net.Listener         //服务器请求监听器
	OnSessionPacket func(*SessionPacket) //处理网络包
	Address         string               //服务器监听端口
	//network  string
}

/*
NewServer
@Description: 创建服务对象
@param address
@param network
@return *Server
*/
func NewServer(address string) *Server {
	s := &Server{
		Address: address,
	}
	return s
}

/*
Run
@Description: 启动服务器
@receiver s
*/
func (s *Server) Run() {

	resolveTCPAddr, err := net.ResolveTCPAddr("tcp6", s.Address)
	if err != nil {
		panic(err)
	}

	tcpListener, err := net.ListenTCP("tcp6", resolveTCPAddr)
	if err != nil {
		panic(err)
	}

	s.listener = tcpListener

	for {
		//循环监听
		conn, err := s.listener.Accept()
		fmt.Println("获取连接")
		if err != nil {
			if _, ok := err.(net.Error); ok {
				fmt.Println(err)
				continue
			}
		}

		//生成session
		newSession := NewSession(conn)
		SessionMgrInstance.AddSession(newSession)
		//启动用户会话，由于内部使用协程启动了读写，所以没必要go出去
		newSession.Run()
		SessionMgrInstance.DelSession(newSession.UId)
	}

}

//获取tcp连接地址
//resolveTCPAddr, err := net.ResolveTCPAddr("tcp6", address)
//if err != nil {
//	panic(err)
//}
//
////获取连接监听器
//tcpListener, err := net.ListenTCP("tcp6", resolveTCPAddr)
//if err != nil {
//	panic(err)
//}
//
//s := &Server{}
//s.listener = tcpListener

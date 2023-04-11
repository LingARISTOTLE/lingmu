package network

import (
	"fmt"
	"net"
)

type Server struct {
	listener        net.Listener
	OnSessionPacket func(*SessionPacket)
	//address  string
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
	//获取tcp连接地址
	resolveTCPAddr, err := net.ResolveTCPAddr("tcp6", address)
	if err != nil {
		panic(err)
	}

	//获取连接监听器
	tcpListener, err := net.ListenTCP("tcp6", resolveTCPAddr)
	if err != nil {
		panic(err)
	}

	s := &Server{}
	s.listener = tcpListener

	return s
}

/*
Run
@Description: 启动服务器
@receiver s
*/
func (s *Server) Run() {

	for {
		//循环监听
		conn, err := s.listener.Accept()
		fmt.Println("获取连接")
		if err != nil {
			continue
		}

		go func() {
			//生成session
			newSession := NewSession(conn)
			SessionMgrInstance.AddSession(newSession)
			//启动用户会话
			newSession.Run()
			SessionMgrInstance.DelSession(newSession.UId)
		}()
	}

}

package network

import (
	"fmt"
	"net"
)

type Server struct {
	listener net.Listener
	address  string
	network  string
}

/*
NewServer
@Description: 创建服务对象
@param address
@param network
@return *Server
*/
func NewServer(address string, network string) *Server {
	return &Server{
		listener: nil,
		address:  address,
		network:  network,
	}
}

func (s *Server) Run() {
	//获取tcp连接地址
	resolveTCPAddr, err := net.ResolveTCPAddr("tcp6", s.address)
	if err != nil {
		fmt.Println(err)
		return
	}
	tcpListener, err := net.ListenTCP("tcp6", resolveTCPAddr)
	if err != nil {
		fmt.Println(err)
		return
	}
	s.listener = tcpListener

	for {
		//循环监听
		conn, err := s.listener.Accept()
		if err != nil {
			continue
		}

		go func() {
			//生成session
			newSession := NewSession(conn)
			newSession.Run()
		}()
	}

}

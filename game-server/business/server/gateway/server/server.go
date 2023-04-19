package server

import "lingmu/game-server/network"

/*
Server
@Description: 服务：维护服务端连接
*/
type Server struct {
	real *network.Server
}

func (s *Server) KickUser() {

}

func (s *Server) KickAllUser() {

}

/*
TransMessageToGateway
@Description: 将消息传递到网关
@receiver s
*/
func (s *Server) TransMessageToGateway() {

}

/*
UpdateRegister
@Description: 更新注册
@receiver s
*/
func (s *Server) UpdateRegister() {

}

/*
CheckRegister
@Description: 检测注册
@receiver s
@return bool
*/
func (s *Server) CheckRegister() bool {
	return true
}

func (s *Server) ReLoginAll() {

}

/*
Loop
@Description: 轮询事件
@receiver s
*/
func (s *Server) Loop() {

	for {
		select {
		//case Message

		}
	}
}

//clients map[int64]*network.Client

//func (s *Server) AddClient(client *network.Client) {
//	s.clients[client.ConnID] = client
//}
//
//func (s *Server) DelClient(client *network.Client) {
//	delete(s.clients, client.ConnID)
//}
//
///*
//GetServerCount
//@Description: 获取服务端数量
//@receiver s
//@return int32
//*/
//func (s *Server) GetServerCount() int32 {
//	return int32(len(s.clients))
//}

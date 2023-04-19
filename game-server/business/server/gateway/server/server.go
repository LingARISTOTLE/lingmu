package server

import "lingmu/game-server/network"

/*
Server
@Description: 服务：维护服务端连接
*/
type Server struct {
	clients map[int64]*network.Client
}

func (s *Server) AddClient(client *network.Client) {
	s.clients[client.ConnID] = client
}

func (s *Server) DelClient(client *network.Client) {
	delete(s.clients, client.ConnID)
}

/*
GetServerCount
@Description: 获取服务端数量
@receiver s
@return int32
*/
func (s *Server) GetServerCount() int32 {
	return int32(len(s.clients))
}

func (s *Server) Loop() {

	for {
		select {
		//case Message

		}
	}
}

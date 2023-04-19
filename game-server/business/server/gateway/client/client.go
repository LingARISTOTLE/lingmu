package client

import "lingmu/game-server/network"

/*
Server
@Description: 服务：维护客户端连接
*/
type Server struct {
	clients map[int64]*network.Client // 维护客户端连接
}

func (s *Server) AddClient(client *network.Client) {
	s.clients[client.ConnID] = client
}

func (s *Server) DelClient(client *network.Client) {
	delete(s.clients, client.ConnID)
}

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

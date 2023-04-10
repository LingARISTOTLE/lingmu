package world

import (
	"lingmu/game-server/manager"
	"lingmu/game-server/network"
)

/*
ManagerHost
@Description: 复制管理所有manager的manager
*/
type ManagerHost struct {
	Pm              *manager.PlayManager
	Server          *network.Server
	Handlers        map[uint64]func(message *network.SessionPacket)
	chSessionPacket chan *network.SessionPacket
}

func NewManagerHost() *ManagerHost {
	m := &ManagerHost{
		Pm: &manager.PlayManager{},
	}
	m.Server = network.NewServer(":8023")
	m.Server.OnSessionPacket = m.OnSessionPacket
	return m
}

var MM *ManagerHost

func (m *ManagerHost) Run() {
	go m.Server.Run()
	go m.Pm.Run()
}

func (m *ManagerHost) OnSessionPacket(packet *network.SessionPacket) {
	if handler, ok := m.Handlers[packet.Msg.Id]; ok {
		handler(packet)
	}
}

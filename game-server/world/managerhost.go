package world

import (
	"lingmu/game-server/manager"
	"lingmu/game-server/network"
	"lingmu/game-server/network/protocol/gen/messageId"
)

/*
ManagerHost
@Description: 复制管理所有manager的manager
*/
type ManagerHost struct {
	Pm              *manager.PlayManager                                         //玩家管理器
	Server          *network.Server                                              //服务器
	Handlers        map[messageId.MessageId]func(message *network.SessionPacket) //消息处理器集合
	chSessionPacket chan *network.SessionPacket                                  //会话包
}

func NewManagerHost() *ManagerHost {
	m := &ManagerHost{
		Pm: &manager.PlayManager{},
	}
	m.Server = network.NewServer(":8023")
	m.Server.OnSessionPacket = m.OnSessionPacket
	m.Handlers = make(map[messageId.MessageId]func(message *network.SessionPacket))
	return m
}

var MM *ManagerHost

func (m *ManagerHost) Run() {
	//注册组管理器处理器
	m.HandlerRegister()
	//启动服务器
	go m.Server.Run()
	//启动玩家管理器
	//go m.Pm.Run()
}

/*
OnSessionPacket
@Description:分发session包处理：调用自己处理器处理，然后根据会话id传递给玩家协程
@receiver m
@param packet
*/
func (m *ManagerHost) OnSessionPacket(packet *network.SessionPacket) {
	//如果是组管理器处理的网络包，那么处理，否则发送给其他管理器，
	if handler, ok := m.Handlers[messageId.MessageId(packet.Msg.Id)]; ok {
		//根据网络包id获得对应的处理方法
		handler(packet)
		return
	}

	//将网络包发送给个人玩家
	if p := m.Pm.GetPlayer(packet.Sess.UId); p != nil {
		p.HandlerParamCh <- packet.Msg
	}
}

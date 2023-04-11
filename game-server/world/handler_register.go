package world

import "lingmu/game-server/network/protocol/gen/messageId"

/*
HandlerRegister
@Description: 组管理器网络包处理器注册
@receiver m
*/
func (m *ManagerHost) HandlerRegister() {
	//m.Handlers[1] = m.UserLogin
	m.Handlers[messageId.MessageId_CSCreatePlayer] = m.CreatePlayer
	m.Handlers[messageId.MessageId_CSLogin] = m.UserLogin
}

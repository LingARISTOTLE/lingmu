package world

import "lingmu/game-server/network/protocol/gen/messageId"

/*
HandlerRegister
@Description: 组管理器网络包处理器注册
@receiver w
*/
func (w *World) HandlerRegister() {
	//w.Handlers[1] = w.UserLogin
	w.Handlers[messageId.MessageId_CSCreatePlayer] = w.CreatePlayer
	w.Handlers[messageId.MessageId_CSLogin] = w.UserLogin
}

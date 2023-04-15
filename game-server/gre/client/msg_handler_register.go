package main

import "lingmu/game-server/network/protocol/gen/messageId"

/*
MessageHandlerRegister
@Description: 注册消息处理器
@receiver c
*/
func (c *ClientManager) MessageHandlerRegister() {
	c.messageHandlers[messageId.MessageId_SCLogin] = c.OnLoginRsp
	c.messageHandlers[messageId.MessageId_SCAddFriend] = c.OnAddFriendRsp
	c.messageHandlers[messageId.MessageId_SCDelFriend] = c.OnDelFriendRsp
	c.messageHandlers[messageId.MessageId_SCSendChatMsg] = c.OnSendChatMsgRsp
	c.messageHandlers[messageId.MessageId_SCCreatePlayer] = c.OnCreatePlayerRsp
}

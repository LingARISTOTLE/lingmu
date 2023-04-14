package player

import "lingmu/game-server/network/protocol/gen/messageId"

/*
HandlerRegister
@Description: 玩家处理器注册，用来注册玩家的业务处理方法，到时候只需要输入字符串就能调用其方法
@receiver p
*/
func (p *Player) HandlerRegister() {
	p.handlers[messageId.MessageId_CSAddFriend] = p.AddFriend
	p.handlers[messageId.MessageId_CSDelFriend] = p.DelFriend
	p.handlers[messageId.MessageId_CSSendChatMsg] = p.ResolveChatMsg
}

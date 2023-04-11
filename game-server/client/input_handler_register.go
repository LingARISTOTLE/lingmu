package main

import (
	"github.com/golang/protobuf/proto"
	"lingmu/game-server/network"
	"lingmu/game-server/network/protocol/gen/messageId"
)

/*
InputHandlerRegister
@Description: 注册用户对应请求处理器方法
@receiver c
*/
func (c *Client) InputHandlerRegister() {
	c.inputHandlers[messageId.MessageId_CSLogin.String()] = c.Login
	c.inputHandlers[messageId.MessageId_CSAddFriend.String()] = c.AddFriend
	c.inputHandlers[messageId.MessageId_CSDelFriend.String()] = c.DelFriend
	c.inputHandlers[messageId.MessageId_CSSendChatMsg.String()] = c.SendChatMsg
}

/*
GetMessageIdByCmd
@Description: 通过cmd命令获取messageId
@receiver c
@param cmd
@return messageId.MessageId
*/
func (c *Client) GetMessageIdByCmd(cmd string) messageId.MessageId {
	mid, ok := messageId.MessageId_value[cmd]
	if ok {
		return messageId.MessageId(mid)
	}
	return messageId.MessageId_None
}

/*
Transport
@Description: 转换message的编码格式,并发送到消息写回管道
@receiver c
@param id
@param message
*/
func (c *Client) Transport(id messageId.MessageId, message proto.Message) {
	//获取message的有限格式wire-format编码
	bytes, err := proto.Marshal(message)
	if err != nil {
		return
	}
	//将编码重新转换为message
	c.cli.ChMsg <- &network.Message{
		Id:   uint64(id),
		Data: bytes,
	}
}

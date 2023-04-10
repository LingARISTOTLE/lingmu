package main

import "lingmu/game-server/network"

type MessageHandler func(packet *network.ClientPacket)

type InputHandler func(param *InputParam)

/*
Login
@Description: 登录处理
@receiver c
@param param
*/
func (c *Client) Login(param *InputParam) {

}

/*
OnLoginRsp
@Description: 登录写回
@receiver c
@param packet
*/
func (c *Client) OnLoginRsp(packet *network.ClientPacket) {

}

/*
AddFriend
@Description: 添加好友
@receiver c
@param param
*/
func (c *Client) AddFriend(param *InputParam) {

}

/*
OnAddFriendRsp
@Description: 添加好友写回
@receiver c
@param packet
*/
func (c *Client) OnAddFriendRsp(packet *network.ClientPacket) {

}

/*
DelFriend
@Description: 删除好友
@receiver c
@param param
*/
func (c *Client) DelFriend(param *InputParam) {

}

/*
OnDelFriendRsp
@Description: 删除好友写回
@receiver c
@param param
*/
func (c *Client) OnDelFriendRsp(packet *network.ClientPacket) {

}

/*
SendChatMsg
@Description: 发送信息
@receiver c
@param param
*/
func (c *Client) SendChatMsg(param *InputParam) {

}

/*
OnSendChatMsgRsp
@Description: 发送信息写回
@receiver c
*/
func (c *Client) OnSendChatMsgRsp(packet *network.ClientPacket) {

}

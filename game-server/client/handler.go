// 所有的方法处理类
package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"lingmu/game-server/network"
	"lingmu/game-server/network/protocol/gen/player"
	"strconv"
)

type MessageHandler func(packet *network.ClientPacket)

type InputHandler func(param *InputParam)

/*
CreatePlayer
@Description: 新建玩家
@receiver c
@param param
*/
func (c *Client) CreatePlayer(param *InputParam) {
	//通过命令获取对应处理方法id
	id := c.GetMessageIdByCmd(param.Command)

	if len(param.Param) != 2 {
		return
	}

	msg := &player.CSCreateUser{
		Username: param.Param[0],
		Password: param.Param[1],
	}

	//转换msg
	c.Transport(id, msg)
}

/*
OnCreatePlayerRsp
@Description: 创建角色响应
@receiver c
@param packet
*/
func (c *Client) OnCreatePlayerRsp(packet *network.ClientPacket) {
	fmt.Println("角色创建成功")
}

/*
Login
@Description: 登录处理
@receiver c
@param param
*/
func (c *Client) Login(param *InputParam) {
	id := c.GetMessageIdByCmd(param.Command)
	if len(param.Param) != 2 {
		return
	}

	msg := &player.CSLogin{
		UserName: param.Param[0],
		Password: param.Param[1],
	}

	c.Transport(id, msg)
}

/*
OnLoginRsp
@Description: 登录写回
@receiver c
@param packet
*/
func (c *Client) OnLoginRsp(packet *network.ClientPacket) {
	//写回包暂且为nil
	rsp := &player.SCLogin{}

	err := proto.Unmarshal(packet.Msg.Data, rsp)
	if err != nil {
		return
	}

	fmt.Println("登录成功")
}

/*
AddFriend
@Description: 添加好友
@receiver c
@param param
*/
func (c *Client) AddFriend(param *InputParam) {
	//获取请求处理方法类型id
	id := c.GetMessageIdByCmd(param.Command)

	//如果命令的参数不是一个，或者命令为0空指令，那么直接返回
	if len(param.Param) != 1 || len(param.Param[0]) == 0 {
		return
	}

	parseUint, err := strconv.ParseUint(param.Param[0], 10, 64)
	if err != nil {
		return
	}

	msg := &player.CSAddFriend{UId: parseUint}

	//传输到发送通道
	c.Transport(id, msg)
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
	//获取请求类型
	id := c.GetMessageIdByCmd(param.Command)

	if len(param.Param) != 1 || len(param.Param[0]) == 0 {
		return
	}

	parseUint, err := strconv.ParseUint(param.Param[0], 10, 64)
	if err != nil {
		return
	}

	//生成删除包
	msg := &player.CSDelFriend{UId: parseUint}

	c.Transport(id, msg)
}

/*
OnDelFriendRsp
@Description: 删除好友写回
@receiver c
@param param
*/
func (c *Client) OnDelFriendRsp(packet *network.ClientPacket) {
	fmt.Println("删除成功")
}

/*
SendChatMsg
@Description: 发送信息
@receiver c
@param param
*/
func (c *Client) SendChatMsg(param *InputParam) {
	//获取消息类型
	id := c.GetMessageIdByCmd(param.Command)

	//如果不是[CSSendChatMsg = 100005][消息内容][消息类型]
	if len(param.Param) != 3 {
		return
	}

	parseUint, err := strconv.ParseUint(param.Param[0], 10, 64)
	if err != nil {
		return
	}

	//解析消息类型：群发/私信/频道等
	parseInt32, err := strconv.ParseInt(param.Param[2], 10, 32)
	if err != nil {
		return
	}

	msg := &player.CSSendChatMsg{
		UId: parseUint,
		Msg: &player.ChatMessage{
			Content: param.Param[1],
			Extra:   nil,
		},
		Category: int32(parseInt32),
	}

	c.Transport(id, msg)

}

/*
OnSendChatMsgRsp
@Description: 发送信息写回
@receiver c
*/
func (c *Client) OnSendChatMsgRsp(packet *network.ClientPacket) {
	fmt.Println("发送消息成功")
}

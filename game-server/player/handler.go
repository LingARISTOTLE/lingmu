package player

import (
	"fmt"
	sugar "github.com/LingARISTOTLE/lingotools"
	"github.com/golang/protobuf/proto"
	"lingmu/game-server/network"
	"lingmu/game-server/network/protocol/gen/player"
)

// Handler 定义方法枚举类型，只要参数是packet *network.SessionPacket无返回值的方法都是player.Handler
type Handler func(packet *network.SessionPacket)

/*
AddFriend
@Description：添加好友
@receiver p：当前对象
@param fId：好友id
*/
func (p *Player) AddFriend(packet *network.SessionPacket) {
	request := &player.CSAddFriend{}
	err := proto.Unmarshal(packet.Msg.Data, request)
	if err != nil {
		return
	}
	//判断好友集合中是否以及存在
	if !sugar.CheckInSlice(request.UId, p.FriendList) {
		p.FriendList = append(p.FriendList, request.UId)
	}

}

/*
DelFriend
@Description:删除好友
@receiver p：当前对象
@param fId：好友id
*/
func (p *Player) DelFriend(packet *network.SessionPacket) {
	request := &player.CSDelFriend{}
	err := proto.Unmarshal(packet.Msg.Data, request)
	if err != nil {
		return
	}
	p.FriendList = sugar.DelOneInSlice(request.UId, p.FriendList)
}

/*
ResolveChatMsg
@Description: 解析消息
@receiver p
@param data
*/
func (p *Player) ResolveChatMsg(packet *network.SessionPacket) {
	//解析聊天包
	request := &player.CSSendChatMsg{}
	err := proto.Unmarshal(packet.Msg.Data, request)
	if err != nil {
		return
	}
	//打印聊天内容
	fmt.Println(request.Msg.Content)
	//收到消息，然后转发给客户端
}

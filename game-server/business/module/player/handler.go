package player

import (
	"fmt"
	sugar "github.com/LingARISTOTLE/lingotools"
	"github.com/golang/protobuf/proto"
	"lingmu/game-server/network"
	"lingmu/game-server/network/protocol/gen/messageId"
	"lingmu/game-server/network/protocol/gen/player"
)

// Handler 定义方法枚举类型，只要参数是packet *network.SessionPacket无返回值的方法都是player.Handler
type Handler func(packet *network.Message)

/*
AddFriend
@Description：添加好友
@receiver p：当前对象
@param fId：好友id
*/
func (p *Player) AddFriend(packet *network.Message) {
	request := &player.CSAddFriend{}
	err := proto.Unmarshal(packet.Data, request)
	if err != nil {
		return
	}
	//判断好友集合中是否以及存在
	if !sugar.CheckInSlice(request.UId, p.FriendList) {
		p.FriendList = append(p.FriendList, request.UId)
	}

	p.SendMsg(messageId.MessageId_SCAddFriend, &player.SCSendChatMsg{})
}

/*
DelFriend
@Description:删除好友
@receiver p：当前对象
@param fId：好友id
*/
func (p *Player) DelFriend(packet *network.Message) {
	request := &player.CSDelFriend{}
	err := proto.Unmarshal(packet.Data, request)
	if err != nil {
		return
	}
	//删除玩家的好友
	p.FriendList = sugar.DelOneInSlice(request.UId, p.FriendList)

	p.SendMsg(messageId.MessageId_SCDelFriend, &player.SCDelFriend{})
}

/*
ResolveChatMsg
@Description: 解析消息
@receiver p
@param data
*/
func (p *Player) ResolveChatMsg(packet *network.Message) {
	//解析聊天包
	request := &player.CSSendChatMsg{}
	err := proto.Unmarshal(packet.Data, request)
	if err != nil {
		return
	}
	//打印聊天内容
	fmt.Println(request.Msg.Content)

	//p.Session.SendMsg(response)
	p.SendMsg(messageId.MessageId_SCSendChatMsg, &player.SCSendChatMsg{})
}

func (p *Player) SendMsg(ID messageId.MessageId, message proto.Message) {
	id := uint64(ID)
	p.Session.AsyncSend(uint16(id), message)
}

//bytes, err := proto.Marshal(&player.SCSendChatMsg{})
//if err != nil {
//	return
//}
//
//response := &network.Message{
//	Id:   uint64(messageId.MessageId_SCAddFriend),
//	Data: bytes,
//}

//p.Session.SendMsg(response)

////生成返回包
//bytes, err := proto.Marshal(&player.SCDelFriend{})
//if err != nil {
//	return
//}
//
//response := &network.Message{
//	Id:   uint64(messageId.MessageId_SCDelFriend),
//	Data: bytes,
//}
//
//p.Session.SendMsg(response)

//收到消息，然后转发给客户端
//bytes, err := proto.Marshal(&player.SCSendChatMsg{})
//if err != nil {
//	return
//}
//
//response := &network.Message{
//	Id:   uint64(messageId.MessageId_SCSendChatMsg),
//	Data: bytes,
//}

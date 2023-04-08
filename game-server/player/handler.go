package player

import (
	"fmt"
	"lingmu/game-server/chat"
	"lingmu/game-server/function"
)

// Handler 定义方法枚举类型，只要参数是interface{}无返回值的方法都是player.Handler
type Handler func(interface{})

/*
AddFriend
@Description：添加好友
@receiver p：当前对象
@param fId：好友id
*/
func (p *Player) AddFriend(data interface{}) {
	fId := data.(uint64)
	if !function.CheckInNumberSlice(fId, p.FriendList) {
		p.FriendList = append(p.FriendList, fId)
	}
}

/*
DelFriend
@Description:删除好友
@receiver p：当前对象
@param fId：好友id
*/
func (p *Player) DelFriend(data interface{}) {
	fId := data.(uint64)
	p.FriendList = function.DelEleInSlice(fId, p.FriendList)
}

/*
ResolveChatMsg
@Description: 解析消息
@receiver p
@param data
*/
func (p *Player) ResolveChatMsg(data interface{}) {
	chatMsg := data.(chat.Message)
	fmt.Println(chatMsg)
	//收到消息，然后转发给客户端
}

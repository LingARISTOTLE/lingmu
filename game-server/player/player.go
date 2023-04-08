package player

import (
	"lingmu/game-server/chat"
	"lingmu/game-server/function"
)

type Player struct {
	UId        uint64
	FriendList []uint64          //朋友
	chChat     chan chat.Private //私聊
}

/*
NewPlayer 构造方法
*/
func NewPlayer() *Player {
	p := &Player{
		UId:        0,
		FriendList: nil,
	}
	return p
}

/*
AddFriend
@Description：添加好友
@receiver p：当前对象
@param fId：好友id
*/
func (p *Player) AddFriend(fId uint64) {
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
func (p *Player) DelFriend(fId uint64) {
	p.FriendList = function.DelEleInSlice(fId, p.FriendList)
}

/*
Run
@Description: 用户协程
@receiver p
*/
func (p *Player) Run() {
	for {
		select {
		case chatMsg := <-p.chChat: //如果chChat管道有值了，那么读取出来处理私信消息
			p.ResolveChatMsg(chatMsg)
		}
	}
}

func (p *Player) ResolveChatMsg(chatMsg chat.Message) {

}

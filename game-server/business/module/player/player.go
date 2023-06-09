package player

import (
	"lingmu/game-server/network"
	"lingmu/game-server/network/protocol/gen/messageId"
)

type Player struct {
	UId            uint64
	FriendList     []uint64                        //朋友
	HandlerParamCh chan *network.Message           //事件通道
	handlers       map[messageId.MessageId]Handler //注册处理方法
	Session        *network.TcpConnX               //用户会话
}

/*
NewPlayer 构造方法
*/
func NewPlayer() *Player {
	p := &Player{
		UId:        0,
		FriendList: make([]uint64, 100),
		handlers:   make(map[messageId.MessageId]Handler),
	}
	//p.HandlerRegister() //将自己的三个方法注册到处理方法中
	return p
}

func (p *Player) Stop() {

}

/*
Start
@Description: 用户协程
@receiver p
*/
func (p *Player) Start() {
	for {
		select {
		case handlerParam := <-p.HandlerParamCh: //循环监听事件，当有事件发生是处理数据
			if fn, ok := p.handlers[messageId.MessageId(handlerParam.Id)]; ok {
				fn(handlerParam)
			}
		}
	}
}

func (p *Player) OnLogin() {
	//从db加载数据初始化
	//同步数据给客户端
}

func (p *Player) OnLogout() {
	//存db
}

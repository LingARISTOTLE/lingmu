package player

import (
	"lingmu/game-server/define"
)

type Player struct {
	UId            uint64
	FriendList     []uint64                 //朋友
	HandlerParamCh chan define.HandlerParam //事件通道
	handlers       map[string]Handler       //注册处理方法
}

/*
NewPlayer 构造方法
*/
func NewPlayer() *Player {
	p := &Player{
		UId:        0,
		FriendList: make([]uint64, 100),
		handlers:   make(map[string]Handler),
	}
	p.HandlerRegister() //将自己的三个方法注册到处理方法中
	return p
}

/*
Run
@Description: 用户协程
@receiver p
*/
func (p *Player) Run() {
	for {
		select {
		case handlerParam := <-p.HandlerParamCh: //循环监听事件，当有事件发生是处理数据
			if fn, ok := p.handlers[handlerParam.HandlerKey]; ok {
				fn(handlerParam.Data)
			}
		}
	}
}

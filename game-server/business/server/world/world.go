package world

import (
	"fmt"
	"lingmu/game-server/aop/logger"
	"lingmu/game-server/business/module/player"
	"lingmu/game-server/network"
	"lingmu/game-server/network/protocol/gen/messageId"
	"os"
	"syscall"
)

/*
World
@Description: 复制管理所有manager的manager
*/
type World struct {
	Pm              *player.Manager                                       //玩家管理器
	Server          *network.Server                                       //服务器
	Handlers        map[messageId.MessageId]func(message *network.Packet) //消息处理器集合
	chSessionPacket chan *network.Packet                                  //会话包
}

func NewWorld() *World {
	w := &World{
		Pm: player.NewPlayerMgr(),
	}
	w.Server = network.NewServer(":8023", 100, 200, logger.Logger)
	w.Server.MessageHandler = w.OnSessionPacket
	w.Handlers = make(map[messageId.MessageId]func(message *network.Packet))
	return w
}

var Oasis *World

func (w *World) Start() {
	//注册组管理器处理器
	w.HandlerRegister()
	//启动服务器
	go w.Server.Run()
	//启动玩家管理器
	go w.Pm.Run()
}

func (w *World) Stop() {

}

/*
OnSessionPacket
@Description:分发session包处理：调用自己处理器处理，然后根据会话id传递给玩家协程
@receiver w
@param packet
*/
func (w *World) OnSessionPacket(packet *network.Packet) {
	//如果是组管理器处理的网络包，那么处理，否则发送给其他管理器，
	if handler, ok := w.Handlers[messageId.MessageId(packet.Msg.Id)]; ok {
		//根据网络包id获得对应的处理方法
		handler(packet)
		return
	}

	//将网络包发送给个人玩家
	if p := w.Pm.GetPlayer(uint64(packet.Conn.ConnID)); p != nil {
		p.HandlerParamCh <- packet.Msg
	}
}

func (w *World) OnSystemSignal(signal os.Signal) bool {
	fmt.Println("接收信号")
	tag := true

	switch signal {
	case syscall.SIGHUP:
	case syscall.SIGPIPE:
	default:
		fmt.Println("等待接收其他信号")
		tag = false
	}

	return tag

}

package world

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"lingmu/game-server/network"
	"lingmu/game-server/network/protocol/gen/player"
	userPlayer "lingmu/game-server/player"
	"time"
)

/*
CreatePlayer
@Description: 创建玩家处理方法
@receiver m
@param message
*/
func (m *ManagerHost) CreatePlayer(message *network.SessionPacket) {
	//创建玩家消息
	msg := &player.CSCreateUser{}
	err := proto.Unmarshal(message.Msg.Data, msg)
	if err != nil {
		return
	}
	fmt.Println("创建玩家", msg)
	m.SendMsg(message.Msg.Id, &player.SCCreateUser{}, message.Sess)
}

/*
SendMsg
@Description: 发送消息处理方法
@receiver m
@param id
@param message
@param session
*/
func (m *ManagerHost) SendMsg(id uint64, message proto.Message, session *network.Session) {
	//将proto的Message解析为[]byte
	bytes, err := proto.Marshal(message)
	if err != nil {
		return
	}
	//生成发送消息的返回包
	rsp := &network.Message{
		Id:   id,
		Data: bytes,
	}
	//调用session发送包
	session.SendMsg(rsp)
}

/*
UserLogin
@Description: 用户登录后包装用户
@receiver m
@param packet
*/
func (m *ManagerHost) UserLogin(packet *network.SessionPacket) {
	msg := &player.CSLogin{}
	err := proto.Unmarshal(packet.Msg.Data, msg)

	if err != nil {
		return
	}

	//创建玩家
	newPlayer := userPlayer.NewPlayer()
	newPlayer.UId = uint64(time.Now().Unix())
	//每个玩家都拥有写回管道
	newPlayer.HandlerParamCh = packet.Sess.WriteCh
	packet.Sess.IsPlayerOnline = true
	//将当前玩家交给PlayerManager去管理
	m.Pm.Add(newPlayer)
	//启动玩家线程
	newPlayer.Run()

}

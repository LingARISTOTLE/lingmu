package world

import (
	"lingmu/game-server/network"
	"lingmu/game-server/player"
)

/*
UserLogin
@Description: 用户登录后包装用户
@receiver m
@param packet
*/
func (m *ManagerHost) UserLogin(packet *network.SessionPacket) {
	newPlayer := player.NewPlayer()
	newPlayer.UId = 111
	newPlayer.HandlerParamCh = packet.Sess.WriteCh
	packet.Sess.IsPlayerOnline = true
	newPlayer.Run()

}

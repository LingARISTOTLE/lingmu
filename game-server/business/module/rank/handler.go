package rank

import (
	"lingmu/game-server/business/module/player"
	"lingmu/game-server/business/module/register"
	"lingmu/game-server/network"
)

func init() {
	register.Register(222, GetRankList)

}

func GetRankList(player *player.Player, packet *network.Packet) {

}

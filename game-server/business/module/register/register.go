package register

import (
	"lingmu/game-server/business/module/player"
	"lingmu/game-server/network"
)

type Fn func(player *player.Player, packet *network.Packet)

func Register(cmd uint32, fn Fn) {
	//装饰器

	//
}

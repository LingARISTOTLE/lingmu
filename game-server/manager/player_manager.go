package manager

import "lingmu/game-server/player"

/*
PlayManager
@Description: 维护在线玩家
*/
type PlayManager struct {
	players map[uint64]player.Player //所有在线玩家
	addPCh  chan player.Player       //用户管道
}

/*
Add
@Description:新增在线玩家
@receiver pm
@param p
*/
func (pm *PlayManager) Add(p player.Player) {
	pm.players[p.UId] = p
	//添加用户在线后，启动用户协程
	go p.Run()
}

func (pm *PlayManager) Del(p player.Player) {
	delete(pm.players, p.UId)
}

/*
Run
@Description: 启动游戏管理器
@receiver pm
*/
func (pm *PlayManager) Run() {
	for {
		select {
		case p := <-pm.addPCh: //当有管道读取到用户时，将其添加到在线集合
			pm.Add(p)
		}
	}
}

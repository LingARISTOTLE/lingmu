package world

import "lingmu/game-server/manager"

/*
ManagerHost
@Description: 复制管理所有manager的manager
*/
type ManagerHost struct {
	Pm manager.PlayManager
}

func NewManagerHost() *ManagerHost {
	m := &ManagerHost{
		Pm: manager.PlayManager{},
	}
	return m
}

var MM *ManagerHost

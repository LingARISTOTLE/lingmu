package activity

import (
	"lingmu/game-server/business/module/base"
	"sync"
)

/*
Manager
@Description: 活动管理器
*/
type Manager struct {
	*base.MetricsBase  //指标类
	*base.DBActionBase //DB操作类
}

var (
	instance *Manager
	onceInit sync.Once
)

func GetMe() *Manager {
	onceInit.Do(func() {
		instance = &Manager{}
	})
	return instance
}

func (a *Manager) OnStart() {
	//TODO implement me
	panic("implement me")
}

func (a *Manager) AfterStart() {
	//TODO implement me
	panic("implement me")
}

func (a *Manager) OnStop() {
	//TODO implement me
	panic("implement me")
}

func (a *Manager) AfterStop() {
	//TODO implement me
	panic("implement me")
}

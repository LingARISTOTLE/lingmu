package activity

import (
	"lingmu/game-server/business/module/base"
	"sync"
)

/*
ConfigManager
@Description: 配置管理器
*/
type ConfigManager struct {
	base.ConfigManagerBase
	Configs sync.Map //策划配置（并发map）

}

/*
Get
@Description: 获取指定id的配置
@receiver m
@param id
@return interface{}
*/
func (m *ConfigManager) Get(id uint32) interface{} {
	var ret any
	m.Configs.Range(func(key, value any) bool {
		idAssert := key.(uint32)
		if idAssert == id {
			ret = value
			return false
		}
		return true
	})
	return ret
}

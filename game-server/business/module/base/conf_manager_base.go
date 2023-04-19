package base

/*
ConfigManagerBase
@Description: 配置管理基础类
*/
type ConfigManagerBase struct {
}

/*
Load
@Description: 加载配置
@receiver c
*/
func (c *ConfigManagerBase) Load() {

}

/*
Get
@Description: 获取配置
@receiver c
@param id
@return interface{}
*/
func (c *ConfigManagerBase) Get(id uint32) interface{} {
	return nil
}

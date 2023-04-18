package template

/*
Conf
@Description: 活动配置
*/
type Conf struct {
	Id          uint32 //id
	Description string //描述
	StartTime   string //开启时间
	EndTime     string //结束时间
	Reward      string //奖励
	Category    string //类别
	Param1      string
	Param2      string
	Param3      string
}

/*
Verify
@Description: 字段检测合法
@receiver c
*/
func (c *Conf) Verify() {
	//todo 一般性检查
	//todo 关联性检查
}

/*
AfterAllVerify
@Description: 字段检测合法
@receiver c
*/
func (c *Conf) AfterAllVerify() {
	//todo 业务逻辑相关检查,可以在所有配置确定都加载好了，再执行此方法
}

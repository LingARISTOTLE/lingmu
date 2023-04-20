package module

/*
MgrInterface
@Description: 管理器接口定义
*/
type MgrInterface interface {
	OnStart()    //启动执行
	AfterStart() //开启后执行
	OnStop()     //关闭执行
	AfterStop()  //关闭后执行
}

/*
Metrics
@Description: 指标接口
*/
type Metrics interface {
	GetName() string     //获取名
	SetName(str string)  //设置名
	Description() string //描述
}

/*
Abstract
@Description: 抽象
*/
type Abstract interface {
}

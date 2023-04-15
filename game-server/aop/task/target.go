package task

/*
Target
@Description: 事件目标
*/
type Target interface {
	CheckDone() bool     //是否完成
	OnNotify(Event)      //执行事件通知
	GetTargetId() uint32 //获取任务id
	SetTaskCB(func())    //设置任务本体
}

/*
TargetBase
@Description: 任务本体
*/
type TargetBase struct {
	TaskCB func() //任务回调
}

func NewTargetBase() *TargetBase {
	return &TargetBase{}
}

func (t *TargetBase) CheckDone() bool {
	return false
}

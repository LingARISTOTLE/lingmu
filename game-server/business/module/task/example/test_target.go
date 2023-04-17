package example

import (
	"lingmu/game-server/business/module/condition"
	"lingmu/game-server/business/module/condition/event"
)

type TTarget struct {
	Id              uint32
	Data            int
	Done            bool
	*condition.Base //任务本体
}

func NewTTarget() *TTarget {
	tt := &TTarget{
		Id:   0,
		Data: 0,
		Done: false,
		Base: condition.NewTargetBase(),
	}
	return tt
}

/*
CheckArrived
@Description: 任务是否被做
@receiver T
@return bool
*/
func (T *TTarget) CheckArrived() bool {
	return T.Done
}

/*
OnNotify
@Description: 事件发布后处理
@receiver T
@param event
*/
func (T *TTarget) OnNotify(event event.Event) {
	e := event.(*TEvent)
	if e.Data == T.Data {
		T.Done = true
	}
	if T.Done {
		//调用回调函数处理
		T.Cb()
	}
}

/*
GetId
@Description: 获取目标id
@receiver T
@return uint32
*/
func (T *TTarget) GetId() uint32 {
	return T.Id
}

func (T *TTarget) SetTaskCB(fn func()) {

}

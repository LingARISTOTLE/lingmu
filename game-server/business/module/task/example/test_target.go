package example

import (
	task2 "lingmu/game-server/business/module/task"
)

type TTarget struct {
	Id                uint32
	Data              int
	Done              bool
	*task2.TargetBase //任务本体
}

func NewTTarget() *TTarget {
	tt := &TTarget{
		Id:         0,
		Data:       0,
		Done:       false,
		TargetBase: task2.NewTargetBase(),
	}
	return tt
}

func (T *TTarget) CheckDone() bool {
	return T.Done
}

/*
OnNotify
@Description: 事件发布后处理
@receiver T
@param event
*/
func (T *TTarget) OnNotify(event task2.Event) {
	e := event.(*TEvent)
	if e.Data == T.Data {
		T.Done = true
	}
	if T.Done {
		//调用回调函数处理
		T.TaskCB()
	}
}

/*
GetTargetId
@Description: 获取目标id
@receiver T
@return uint32
*/
func (T *TTarget) GetTargetId() uint32 {
	return T.Id
}

func (T *TTarget) SetTaskCB(fn func()) {

}
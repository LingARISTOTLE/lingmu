package example

import (
	"lingmu/game-server/business/module/task"
)

type TEvent struct {
	Data        int
	Subscribers []task.Target //订阅用户
}

func (t *TEvent) Notify() {
	for _, subscriber := range t.Subscribers {
		subscriber.OnNotify(t)
	}
}

/*
Attach
@Description: 添加任务订阅
@receiver t
@param target
*/
func (t *TEvent) Attach(target task.Target) {
	t.Subscribers = append(t.Subscribers, target)
}

func (t *TEvent) Detach(id uint32) {
	for i, subscriber := range t.Subscribers {
		if subscriber.GetTargetId() == id {
			t.Subscribers = append(t.Subscribers, t.Subscribers[i+1:]...)
		}
	}
}

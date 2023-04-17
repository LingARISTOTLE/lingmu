package example

import (
	"lingmu/game-server/business/module/condition"
)

type TEvent struct {
	Data        int
	Subscribers []condition.Condition //订阅用户
}

/*
Notify
@Description: 通知所有订阅者调用 处理通知 方法
@receiver t
*/
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
func (t *TEvent) Attach(target condition.Condition) {
	t.Subscribers = append(t.Subscribers, target)
}

func (t *TEvent) Detach(id uint32) {
	for i, subscriber := range t.Subscribers {
		if subscriber.GetId() == id {
			t.Subscribers = append(t.Subscribers, t.Subscribers[i+1:]...)
		}
	}
}

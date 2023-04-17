package event

import "lingmu/game-server/business/module/condition"

type Event interface {
	Notify()
	Attach(condition condition.Condition)
	Detach(id uint32)
}

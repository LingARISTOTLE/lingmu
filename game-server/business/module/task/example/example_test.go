package example

import (
	"fmt"
	"lingmu/game-server/business/module/condition"
	"testing"
)

func TestEvent(t *testing.T) {
	tEvent := TEvent{
		Subscribers: make([]condition.Condition, 0),
	}

	tg := &TTarget{
		Id:   111,
		Data: 1,
		Done: false,
	}

	tEvent.Attach(tg)
	tEvent.Data = 1
	tEvent.Notify()
	fmt.Println("CheckArrived", tg.CheckArrived())
}

func TestTask(t *testing.T) {
	NewTTask(nil)
}

package example

import (
	"fmt"
	"lingmu/game-server/aop/task"
	"testing"
)

func TestName(t *testing.T) {
	tEvent := TEvent{
		Subscribers: make([]task.Target, 0),
	}

	tg := &TTarget{
		Id:   111,
		Data: 1,
		Done: false,
	}

	tEvent.Attach(tg)
	tEvent.Data = 1
	tEvent.Notify()
	fmt.Println("CheckDone", tg.CheckDone())
}

package example

import (
	"lingmu/game-server/business/module/condition"
	task2 "lingmu/game-server/business/module/task"
)

type TTask struct {
	Conf    *task2.Config
	Next    *TTask
	Status  task2.Status
	Targets []condition.Condition
}

func NewTTask(config *task2.Config) *TTask {
	t := &TTask{
		Conf: config,
	}
	return t
}

func (t *TTask) Finish() {
	t.Status = task2.FINISH
}

func (t *TTask) Accept(config *task2.Config) {
	t.Status = task2.FINISH
}

func (t *TTask) TargetDoneCallBack() {
	count := 0
	for _, target := range t.Targets {
		if target.CheckArrived() {
			count++
		}
	}
	if count == len(t.Targets) {
		t.Finish()
	}
}

package task

type Task struct {
	Conf    *Config
	Next    *Task
	Status  Status
	Targets []*Target
}

func NewTask(config *Config) *Task {
	t := &Task{
		Conf: config,
	}
	return t
}

/*
Accept
@Description: 接收任务
@receiver t
@param config
*/
func (t *Task) Accept(config *Config) {
	t.Status = ACCEPT
}

/*
Finish
@Description: 完成任务
@receiver t
*/
func (t *Task) Finish() {
	t.Status = FINISH
}

/*
TargetDoneCallBack
@Description: 任务完成后回调
@receiver t
*/
func (t *Task) TargetDoneCallBack() {

}

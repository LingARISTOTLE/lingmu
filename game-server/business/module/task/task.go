package task

type Task interface {
	Accept(conf *Config) //接收任务
	Finish()             //完成任务
	TargetDoneCallBack() //任务完成后回调
}

type Base struct {
}

/*
Accept
@Description: 接收任务
@receiver t
@param config
*/
func (b *Base) Accept(config *Config) {

}

/*
Finish
@Description: 完成任务
@receiver t
*/
func (b *Base) Finish() {

}

/*
TargetDoneCallBack
@Description: 任务完成后回调
@receiver t
*/
func (b *Base) TargetDoneCallBack() {

}

//	func NewTask(config *Config) *Task {
//		t := &Task{
//			Conf: config,
//		}
//		return t
//	}
//	Next
//*Task
//Status  Status
//Targets []*Target

package activity

/*
Activity
@Description: 活动接口
*/
type Activity interface {
	CheckInTimeRange() bool // 检测是否在活动事件范围内
}

///*
//Conf
//@Description: 活动
//*/
//type Conf struct {
//	Id uint32
//}

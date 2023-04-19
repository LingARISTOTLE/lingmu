package template

import "time"

/*
HappyNewYear
@Description: 新年活动
*/
type HappyNewYear struct {
	Id        uint32    //活动id
	StartTime time.Time //开始时间
	EndTime   time.Time //结束时间
}

func (h *HappyNewYear) Init(conf Conf) *HappyNewYear {
	return &HappyNewYear{}
}

/*
GetReward
@Description: 获取奖励
@receiver h
*/
func (h *HappyNewYear) GetReward() {

}

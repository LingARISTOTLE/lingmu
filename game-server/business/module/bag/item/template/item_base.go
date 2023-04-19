package template

import "time"

/*
ItemBase
@Description: 物品实体
*/
type ItemBase struct {
	Id             uint32 `json:"id"`             //物品id
	Num            int64  `json:"num"`            //物品数量
	LastChangeTime int64  `json:"lastChangeTime"` //最后修改时间
	UseTime        int64  `json:"useTime"`        //使用时间
}

func (i *ItemBase) Add(delta int64) {
	i.Num += delta
	i.LastChangeTime = time.Now().Unix()
}

func (i *ItemBase) Delete(delta int64) {
	i.Num -= delta
	i.LastChangeTime = time.Now().Unix()
}

func (i *ItemBase) GetNum() int64 {
	return i.Num
}

func (i *ItemBase) GetId() uint32 {
	return i.Id
}

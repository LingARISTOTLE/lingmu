package instance

import (
	"lingmu/game-server/business/module/bag/item"
	"lingmu/game-server/business/module/bag/item/template"
	"sync"
)

/*
NormalBag
@Description: 普通背包
*/
type NormalBag struct {
	Data sync.Map
}

func (n NormalBag) AddItem(item item.Item) {
	value, ok := n.Data.Load(item.GetId())
	if ok {
		value.(*template.ItemBase).Add(item.GetNum())
	}
}

func (n *NormalBag) DelItem(item item.Item) {
	value, ok := n.Data.Load(item.GetId())
	if ok {
		value.(*template.ItemBase).Delete(item.GetNum())
	}
}

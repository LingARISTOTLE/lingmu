package bag

import "lingmu/game-server/business/module/bag/item"

/*
Bag
@Description: 背包
*/
type Bag interface {
	AddItem(item item.Item) //添加物品
	DelItem(item item.Item) //删除物品
}

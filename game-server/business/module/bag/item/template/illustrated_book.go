package template

type IllustratedBookType int

const (
	Fish   IllustratedBookType = iota + 1 //鱼
	Flower                     = 2        //花
)

/*
IllustratedBook
@Description: 物品说明书
*/
type IllustratedBook struct {
	*ItemBase
	Category IllustratedBookType //物品类型
}

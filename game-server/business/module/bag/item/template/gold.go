package template

type GoldType int

const (
	NormalGold GoldType = iota + 1 //黄金
	Diamond             = 2        //砖石
)

/*
Gold
@Description: 武器品质
*/
type Gold struct {
	Category GoldType `json:"category"`
	*ItemBase
}

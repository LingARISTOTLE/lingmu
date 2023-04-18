package random_weight

import (
	"errors"
	"math/rand"
	"sort"
)

/*
Choice
@Description: 奖项
*/
type Choice struct {
	Item   interface{} //实体
	Weight uint        //权重
}

func NewChoice(item interface{}, weight uint) Choice {
	return Choice{
		Item:   item,
		Weight: weight,
	}
}

/*
Chooser
@Description: 卡池
*/
type Chooser struct {
	data   []Choice //奖励池
	totals []int    //每个物品所在的权重范围
	max    int      //最大权重
}

var (
	errWeightOverflow = errors.New("sum of Choice Weights exceeds max int")
	errNoValidChoices = errors.New("zero Choices with Weight >= 1")
)

const (
	intSize = 32 << (^uint(0) >> 63) // cf. strconv.IntSize
	maxInt  = 1<<(intSize-1) - 1
)

func NewChooser(choices ...Choice) (*Chooser, error) {
	//按照权重排序奖项
	sort.Slice(choices, func(i, j int) bool {
		return choices[i].Weight < choices[j].Weight
	})

	//为每个奖项标记
	totals := make([]int, len(choices))
	//总权重
	runningTotal := 0

	for i, c := range choices {
		weight := int(c.Weight)

		if (maxInt - runningTotal) <= weight {
			//剩余权重小于需要权重
			return nil, errWeightOverflow
		}

		//累计总权重
		runningTotal += weight
		totals[i] = runningTotal
	}

	//没有权重
	if runningTotal < 1 {
		return nil, errNoValidChoices
	}

	return &Chooser{
		data:   choices,
		totals: totals,
		max:    runningTotal,
	}, nil
}

/*
Pick
@Description: 随机获取奖励
@receiver c
@return interface{}
*/
func (c Chooser) Pick() interface{} {
	//取总权重中的任一个数
	r := rand.Intn(c.max) + 1
	//判断该数在那个一个物品的权重范围内
	i := searchInts(c.totals, r)
	return c.data[i].Item
}

/*
PickSource
@Description: 随机获取奖励
@receiver c
@param rs
@return interface{}
*/
func (c Chooser) PickSource(rs *rand.Rand) interface{} {
	r := rs.Intn(c.max) + 1
	i := searchInts(c.totals, r)
	return c.data[i].Item
}

/*
searchInts
@Description: 二分查找当前随机数命中的奖项
@param a
@param x
@return int
*/
func searchInts(a []int, x int) int {
	i, j := 0, len(a)
	//二分查找
	for i < j {
		h := int(uint(i+j) >> 1) // avoid overflow when computing h
		if a[h] < x {
			i = h + 1
		} else {
			j = h
		}
	}

	return i
}

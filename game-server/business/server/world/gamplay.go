package world

import (
	"lingmu/game-server/business/module/activity"
	"lingmu/game-server/business/module/bag"
	"lingmu/game-server/business/module/chat"
	"lingmu/game-server/business/module/email"
	"lingmu/game-server/business/module/friend"
	"lingmu/game-server/business/module/minigame"
	"lingmu/game-server/business/module/rank"
	"lingmu/game-server/business/module/recharge"
	"lingmu/game-server/business/module/task"
)

/*
GamePlay
@Description: 系统
*/
type GamePlay struct {
	activity activity.Abstract //活动
	bag      bag.Abstract      //背包
	chat     chat.Abstract     //聊天
	rank     rank.Abstract     //排名
	email    email.Abstract    //邮箱
	friend   friend.Abstract   //好友
	minigame minigame.Abstract //小游戏
	recharge recharge.Abstract //充值
	task     task.Abstract     //任务
}
type Option func(play *GamePlay) *GamePlay

/*
WithActivity
@Description: 接受活动
@param activity
@return Option
*/
func WithActivity(activity activity.Abstract) Option {
	return func(play *GamePlay) *GamePlay {
		play.activity = activity
		return play
	}
}

func NewGamePlay(option ...Option) *GamePlay {
	g := &GamePlay{}
	for _, op := range option {
		op(g)
	}
	return g
}

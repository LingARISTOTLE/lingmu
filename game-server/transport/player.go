package transport

/*
Player
@Description: 玩家实体对象
*/
type Player struct {
	Uid        uint64   `bson:"uid"`
	NickName   string   `bson:"nickName"`
	Sex        int      `bson:"sex"`
	FriendList []uint64 `bson:"friendList"`
}

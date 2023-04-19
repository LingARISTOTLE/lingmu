package friend

/*
Info
@Description: 好友实体
*/
type Info struct {
	UId      uint64
	ChatTime int64  //聊天时间
	AddTime  int64  //添加时间
	Tag      string //备注
}

/*
Request
@Description: 好友请求
*/
type Request struct {
	Userid  uint64 // 玩家ID
	OpTime  int64  // 操作时间
	AddType int32  // 申请加好友的途径
}

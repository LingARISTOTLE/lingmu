package network

import "net"

/*
SessionPacket
@Description: 服务端包
*/
type SessionPacket struct {
	Msg  *Message //服务端网络包
	Sess *Session //会话（包含了Conn）
}

/*
ClientPacker
@Description: 客户端包
*/
type ClientPacker struct {
	Msg  *Message  //服务端网络包
	Conn *net.Conn //会话（包含了Conn）
}

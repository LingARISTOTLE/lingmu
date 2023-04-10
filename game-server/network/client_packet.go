package network

import "net"

/*
ClientPacket
@Description: 客户端网络包
*/
type ClientPacket struct {
	Msg  *Message //传输的信息
	Conn net.Conn //连接
}

package network

/*
SessionPacket
@Description: 服务端包
*/
type SessionPacket struct {
	Msg  *Message
	Sess *Session
}

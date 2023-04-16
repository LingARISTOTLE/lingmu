package network

/*
IPacker
@Description: 处理网络包接口
*/
type IPacker interface {
	Pack(MsgId uint16, msg interface{}) ([]byte, error) //网络包打包
	UnPack([]byte) (*Message, error)                    //网络包解包
	Read(x *TcpConnX) ([]byte, error)                   //读取TcpConnX中的数据
}

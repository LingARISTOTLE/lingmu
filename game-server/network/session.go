package network

import (
	"encoding/binary"
	"fmt"
	"lingmu/abandoned"
	"net"
	"time"
)

type Session struct {
	UId            uint64                  //连接序号
	Conn           net.Conn                //连接
	IsClose        bool                    //连接是否关闭
	packer         *abandoned.NormalPacker //打包方式
	WriteCh        chan *Message           //用于写回的通道，当处理完请求后会生成响应包放在chWrite管道
	IsPlayerOnline bool                    //玩家是否在线
	MessageHandler func(packet *SessionPacket)
}

func NewSession(conn net.Conn) *Session {
	return &Session{
		Conn:    conn,
		packer:  abandoned.NewNormalPacker(binary.BigEndian), //采用大端方式去解析
		WriteCh: make(chan *Message, 10),                     //新建一个Message管道，一次能进行10个Message的写回
	}
}

// Run 双方建立连接后就会一直有一个Session，所以在Session中启动读和写两个协程
func (s *Session) Run() {
	go s.Read()
	go s.Write()
}

/*
Read
@Description: 读取网络数据
@receiver s
*/
func (s *Session) Read() {
	for {
		//设置超时时间为1s
		err := s.Conn.SetReadDeadline(time.Now().Add(time.Second))
		if err != nil {
			fmt.Println(err)
			continue
		}

		//通过连接拿到网络包
		message, err := s.packer.UnPack(s.Conn)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println("server receive message:", string(message.Data))
		//处理接收到的message
		//生成服务器网络包，处理消息
		s.MessageHandler(&SessionPacket{
			Msg:  message,
			Sess: s,
		})
		//...
		//这里应该将生成的网络包交给组管理器的OnMessagePacket去处理
		//...

		//处理完后写回
		//s.WriteCh <- &Message{
		//	Id:   555,
		//	Data: []byte("服务端成功获取网络包"),
		//}
	}

}

/*
Write
@Description: 写回客户端
@receiver s
*/
func (s *Session) Write() {

	for {
		select {
		case message := <-s.WriteCh: //SessionPacket，那么发送
			s.send(message)
		}
	}

}

/*
send
@Description: 发送消息
@receiver s
@param message
*/
func (s *Session) send(message *Message) {
	//设置网络包发送超时事件
	err := s.Conn.SetWriteDeadline(time.Now().Add(time.Second))
	if err != nil {
		fmt.Println(err)
		return
	}

	bytes, err := s.packer.Pack(message)
	if err != nil {
		fmt.Println(err)
		return
	}

	s.Conn.Write(bytes)
}

/*
SendMsg
@Description: 发送消息到写回队列
@receiver s
@param message
*/
func (s *Session) SendMsg(message *Message) {
	s.WriteCh <- message
}

//设置超时时间
//err := s.Conn.SetWriteDeadline(time.Now().Add(time.Second))
//if err != nil {
//	fmt.Println(err)
//}

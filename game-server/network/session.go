package network

import (
	"encoding/binary"
	"fmt"
	"net"
	"time"
)

type Session struct {
	conn    net.Conn      //连接
	packer  *NormalPacker //打包方式
	chWrite chan *Message //用于写回的通道，当处理完请求后会生成响应包放在chWrite管道
}

func NewSession(conn net.Conn) *Session {
	return &Session{
		conn:    conn,
		packer:  NewNormalPacker(binary.BigEndian), //采用大端方式去解析
		chWrite: make(chan *Message, 1),            //新建一个Message管道，一次只能进行单个Message的写回
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
		//设置超时时间
		err := s.conn.SetReadDeadline(time.Now().Add(time.Second))
		if err != nil {
			fmt.Println(err)
			continue
		}

		//通过连接拿到网络包
		message, err := s.packer.UnPack(s.conn)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println("server receive message:", string(message.Data))
		//处理接收到的message
		//...

		//处理完后写回
		s.chWrite <- &Message{
			Id:   999,
			Data: []byte("connection success!"),
		}
	}

}

func (s *Session) Write() {
	//设置超时时间
	err := s.conn.SetWriteDeadline(time.Now().Add(time.Second))
	if err != nil {
		fmt.Println(err)
	}

	for {
		select {
		case message := <-s.chWrite: //如果管道中有message，那么发送
			s.send(message)
		}
	}

}

func (s *Session) send(message *Message) {
	bytes, err := s.packer.Pack(message)
	if err != nil {
		fmt.Println(err)
	}
	_, err = s.conn.Write(bytes)
	if err != nil {
		fmt.Println(err)
	}
}

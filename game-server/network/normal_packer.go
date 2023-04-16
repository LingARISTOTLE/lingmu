package network

import (
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/golang/protobuf/proto"
	"io"
	"net"
	"time"
)

/*
NormalPacker
@Description: 网络打包方式
*/
type NormalPacker struct {
	ByteOrder binary.ByteOrder //ByteOrder 指定如何将字节slices转换为 16、32 或 64 位无符号整数。
}

func NewNormalPacker(order binary.ByteOrder) *NormalPacker {
	return &NormalPacker{
		ByteOrder: order,
	}
}

/*
Pack 一个pack结构：数据包长度---messageId---data数据
@Description: 将服务器的数据打包，服务器——>网络
@receiver p
@param message
@return []byte
@return error
*/
func (p *NormalPacker) Pack(msgID uint16, msg interface{}) ([]byte, error) {
	//强转为proto.Message类型
	pbMsg, ok := msg.(proto.Message)
	if !ok {
		return []byte{}, fmt.Errorf("msg is not protobuf message")
	}
	//序列化
	data, err := proto.Marshal(pbMsg)
	if err != nil {
		return data, err
	}
	//数据包长度---messageId---data数据
	buffer := make([]byte, 8+8+len(data))
	p.ByteOrder.PutUint64(buffer[0:8], uint64(len(buffer)))
	p.ByteOrder.PutUint64(buffer[8:16], uint64(msgID))
	copy(buffer[16:], data)
	return buffer, nil
}

/*
Unpack
@Description: 将网络包中的数据解包
@receiver p
@param data
@return *Message
@return error
*/
func (p *NormalPacker) Unpack(data []byte) (*Message, error) {
	if len(data) < 2 {
		return nil, errors.New("protobuf data too short")
	}
	msgID := p.ByteOrder.Uint16(data[:2])
	msg := &Message{
		Id:   uint64(msgID),
		Data: data[2:],
	}
	return msg, nil
}

/*
Read
@Description: 读数据
@receiver p
@param conn
@return []byte
@return error
*/
func (p *NormalPacker) Read(conn *TcpConnX) ([]byte, error) {
	err := conn.Conn.(*net.TCPConn).SetReadDeadline(time.Now().Add(time.Second * 10))
	if err != nil {
		return nil, err
	}
	buffer := make([]byte, 8+8)

	if _, err := io.ReadFull(conn.Conn, buffer); err != nil {
		return nil, err
	}
	totalSize := p.ByteOrder.Uint64(buffer[:8])
	dataSize := totalSize - 8 - 8
	data := make([]byte, 8+dataSize)
	copy(data[:8], buffer[8:])
	if _, err := io.ReadFull(conn.Conn, data[8:]); err != nil {
		return nil, err
	}
	return data, nil
}

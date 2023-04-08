package network

import (
	"encoding/binary"
	"io"
	"net"
	"time"
)

type NormalPacker struct {
	Order binary.ByteOrder //ByteOrder 指定如何将字节slices转换为 16、32 或 64 位无符号整数。
}

func NewNormalPacker(order binary.ByteOrder) *NormalPacker {
	return &NormalPacker{
		Order: order,
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
func (p *NormalPacker) Pack(message *Message) ([]byte, error) {
	//创建一个buffer，大小为8字节数据长度（64位），8字节id，若干字节存储数据
	buffer := make([]byte, 8+8+len(message.Data))
	p.Order.PutUint64(buffer[:8], uint64(len(buffer)))
	p.Order.PutUint64(buffer[8:16], message.Id)
	copy(buffer[16:], message.Data)
	return buffer, nil
}

/*
UnPack
@Description: 将网络包中的数据解包
@receiver p
@param reader
*/
func (p *NormalPacker) UnPack(reader io.Reader) (*Message, error) {
	err := reader.(*net.TCPConn).SetReadDeadline(time.Now().Add(time.Second))
	if err != nil {
		return nil, err
	}
	//先读取前16位
	buffer := make([]byte, 8+8)

	_, err = io.ReadFull(reader, buffer)
	if err != nil {
		return nil, err
	}

	totalLen := p.Order.Uint64(buffer[:8])
	id := p.Order.Uint64(buffer[8:])
	//数据包长度=总长度-前16位
	dataLen := totalLen - 16
	dataBuffer := make([]byte, dataLen)
	_, err = io.ReadFull(reader, dataBuffer)
	if err != nil {
		return nil, err
	}

	msg := &Message{
		Id:   id,
		Data: dataBuffer,
	}

	return msg, nil
}

package network

import (
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/golang/protobuf/proto"
	"io"
	"math"
)

/*
BufferPacker
@Description: 网络包缓冲处理类，规定了网络包的规格大小，缓冲大小以及大端小段等
*/
type BufferPacker struct {
	lenMsgLen int32            //包长度
	minMsgLen uint32           //最小长度
	maxMsgLen uint32           //最大长度
	recvBuff  *ByteBuffer      //接收缓冲
	sendBuff  *ByteBuffer      //发送缓冲
	byteOrder binary.ByteOrder //大小端
}

/*
newInActionPacker
@Description: 新建网络包缓冲处理结构体
@return *BufferPacker : 网络包的规格为 [长度][数据内容]
*/
func newInActionPacker() *BufferPacker {
	msgParser := &BufferPacker{
		lenMsgLen: 4,                   //默认长度为第4级别
		minMsgLen: 2,                   //默认最小长度为2B
		maxMsgLen: 2 * 1024 * 1024,     //默认最大长度为2MB
		recvBuff:  NewByteBuffer(),     //生成接收缓冲
		sendBuff:  NewByteBuffer(),     //生成发送缓冲
		byteOrder: binary.LittleEndian, //默认小端
	}
	return msgParser
}

/*
SetMsgLen
@Description: 调整网络包规格
@receiver p
@param lenMsgLen
@param minMsgLen
@param maxMsgLen
*/
func (p *BufferPacker) SetMsgLen(lenMsgLen int32, minMsgLen uint32, maxMsgLen uint32) {
	if lenMsgLen == 1 || lenMsgLen == 2 || lenMsgLen == 4 {
		p.lenMsgLen = lenMsgLen
	}
	if minMsgLen != 0 {
		p.minMsgLen = minMsgLen
	}
	if maxMsgLen != 0 {
		p.maxMsgLen = maxMsgLen
	}

	var max uint32
	switch p.lenMsgLen {
	case 1:
		max = math.MaxUint8
	case 2:
		max = math.MaxUint16
	case 4:
		max = math.MaxUint32
	}
	if p.minMsgLen > max {
		p.minMsgLen = max
	}
	if p.maxMsgLen > max {
		p.maxMsgLen = max
	}
}

/*
Read
@Description: 将conn中的数据读入缓冲，并返回读入数据
@receiver p
@param conn：连接，可以通过连接拿到网络包:[长度][数据]
@return []byte：读取到字节数组
@return error
*/
func (p *BufferPacker) Read(conn *TcpConnX) ([]byte, error) {
	//确保缓冲区能够写下包长度lenMsgLen，不够则扩容
	p.recvBuff.EnsureWritableBytes(p.lenMsgLen)
	//从连接TcpConnX中读取长度为lenMsgLen的数据到接收缓冲中
	//这个readLen就是整个网络包的长度
	readLen, err := io.ReadFull(conn, p.recvBuff.WriteBuff()[:p.lenMsgLen])
	if err != nil {
		return nil, fmt.Errorf("%v readLen:%v", err, readLen)
	}

	//写指针右移readLen，
	p.recvBuff.WriteBytes(int32(readLen))

	//网络包长度
	var msgLen uint32

	//根据网络包前面的[长度]来判断，有多少位被用于长度表示
	switch p.lenMsgLen {
	case 2:
		//读取2字节数据
		msgLen = uint32(p.recvBuff.ReadInt16())
	case 4:
		//读取4字节数据
		msgLen = uint32(p.recvBuff.ReadInt32())
	}

	//判断是否越界
	if msgLen > p.maxMsgLen {
		return nil, errors.New("message too long")
	} else if msgLen < p.minMsgLen {
		return nil, errors.New("message too short")
	}

	//确保能写msgLen字节
	p.recvBuff.EnsureWritableBytes(int32(msgLen))

	//读取全部数据（写指针————数据长度）
	rLen, err := io.ReadFull(conn, p.recvBuff.WriteBuff()[:msgLen])

	if err != nil {
		return nil, fmt.Errorf("%v msgLen:%v readLen:%v", err, msgLen, rLen)
	}

	//写指针右移rLen
	p.recvBuff.WriteBytes(int32(rLen))

	// 保留了2字节flag 暂时未处理
	//跳过2字节保留字段
	p.recvBuff.Skip(2)

	//减去2字节的保留字段长度
	return p.recvBuff.NextBytes(int32(msgLen - 2)), nil
}

/*
Write
@Description: 将数据写入写缓冲
@receiver p
@param conn
@param buff
@return error
*/
func (p *BufferPacker) Write(conn *TcpConnX, buff ...byte) error {
	//获取缓冲长度
	msgLen := uint32(len(buff))

	if msgLen > p.maxMsgLen {
		return errors.New("message too long")
	} else if msgLen < p.minMsgLen {
		return errors.New("message too short")
	}

	//添加[长度]字节表示位
	switch p.lenMsgLen {
	case 2:
		p.sendBuff.AppendInt16(int16(msgLen))
	case 4:
		p.sendBuff.AppendInt32(int32(msgLen))
	}

	//写入缓冲
	p.sendBuff.Append(buff)
	//读取全部写缓冲数据
	writeBuff := p.sendBuff.ReadBuff()[:p.sendBuff.Length()]
	//将写缓冲数据写回
	_, err := conn.Write(writeBuff)
	//重制写缓冲状态
	p.sendBuff.Reset()

	return err

}

/*
reset
@Description: 重置缓冲
@receiver p
*/
func (p *BufferPacker) reset() {
	p.recvBuff = NewByteBuffer()
	p.sendBuff = NewByteBuffer()
}

/*
Pack
@Description: 打包网络包，这里默认[长度]的咱用字节为2
@receiver p
@param msgID
@param msg
@return []byte
@return error
*/
func (p *BufferPacker) Pack(msgID uint16, msg interface{}) ([]byte, error) {
	//强转当前message
	pbMsg, ok := msg.(proto.Message)
	if !ok {
		return []byte{}, fmt.Errorf("message is not protobuf message")
	}

	//序列化
	data, err := proto.Marshal(pbMsg)
	if err != nil {
		return data, err
	}

	buf := make([]byte, 4+len(data))

	//前网络包的前2字节记录长度，2字节后记录msgID（2字节）
	//4byte = len(flag)[2byte] + len(msgID)[2byte]
	if p.byteOrder == binary.LittleEndian {
		binary.LittleEndian.PutUint16(buf[0:2], 0)
		binary.LittleEndian.PutUint16(buf[2:], msgID)
	} else {
		binary.BigEndian.PutUint16(buf[0:2], 0)
		binary.BigEndian.PutUint16(buf[2:], msgID)
	}
	//将网络包数据放入4字节后
	copy(buf[4:], data)
	return buf, err

}

/*
Unpack
@Description:  解包 id + protobuf data，这里默认[长度]的咱用字节为2
@receiver p
@param data
@return *Message
@return error
*/
func (p *BufferPacker) Unpack(data []byte) (*Message, error) {
	if len(data) < 2 {
		return nil, errors.New("protobuf data to short")
	}

	msgID := p.byteOrder.Uint16(data[:2])
	msg := &Message{
		Id:   uint64(msgID),
		Data: data[2:],
	}
	return msg, nil
}

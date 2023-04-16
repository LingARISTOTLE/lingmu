package network

import "encoding/binary"

const (
	cheapPrependSize = 8    //预留开始位置
	initialSize      = 1024 //初始化长度
)

/*
ByteBuffer
@Description: 字节缓冲
*/
type ByteBuffer struct {
	mBuffer             []byte //缓冲大小
	mCapacity           int32  //最大容量
	readIndex           int32  //读指针，记录当前可读的前一位
	writeIndex          int32  //写指针，记录当前可写的前一位
	reservedPrependSize int32  //预留位置，默认从该位置为初始操作位置
	littleEndian        bool   //小端排序
}

/*
NewByteBuffer
@Description: 创建一个字节buffer
@return *ByteBuffer
*/
func NewByteBuffer() *ByteBuffer {
	return &ByteBuffer{
		mBuffer:             make([]byte, cheapPrependSize+initialSize), //缓冲大小
		mCapacity:           cheapPrependSize + initialSize,             //最大容量
		readIndex:           cheapPrependSize,                           //读指针
		writeIndex:          cheapPrependSize,                           //写指针
		reservedPrependSize: cheapPrependSize,                           //预置位置默认从8开始，前面用于特殊记录
		littleEndian:        true,                                       //小端排序
	}
}

/*
SetByteOrder
@Description: 设置大小端
@receiver bf
@param littleEndian
*/
func (bf *ByteBuffer) SetByteOrder(littleEndian bool) {
	bf.littleEndian = littleEndian
}

/*
Length
@Description: 获取到当前未读部分长度
@receiver bf
@return int32
*/
func (bf *ByteBuffer) Length() int32 {
	return bf.writeIndex - bf.readIndex
}

func (bf *ByteBuffer) Swap(other *ByteBuffer) {

}

/*
Skip
@Description: 跳过需要读取的指定长度
@receiver bf
@param len
*/
func (bf *ByteBuffer) Skip(len int32) {
	if len < bf.Length() {
		bf.readIndex = bf.readIndex + len
	} else {
		bf.Reset()
	}
}

/*
Reset
@Description: 重制指针
@receiver bf
*/
func (bf *ByteBuffer) Reset() {
	bf.Truncate(0)

}

func (bf *ByteBuffer) Retrieve(len int32) {
	bf.Skip(len)
}

/*
Truncate
@Description: 回退写指针，回退到读指针后的n位，如果n在写指针前，那么不回退
@receiver bf
@param n
*/
func (bf *ByteBuffer) Truncate(n int32) {
	//如果传入0，那么重制
	if n == 0 {
		bf.readIndex = bf.reservedPrependSize
		bf.writeIndex = bf.reservedPrependSize
	} else if bf.writeIndex > (bf.readIndex + n) {
		bf.writeIndex = bf.readIndex + n
	}
}

/*
Reserve
@Description: 扩容
@receiver bf
@param len
*/
func (bf *ByteBuffer) Reserve(len int32) {
	if bf.mCapacity >= len+bf.reservedPrependSize {
		return
	}
	bf.grow(len + bf.reservedPrependSize)
}

/*
Append
@Description: 在原字节数组后追加字节数组
@receiver bf
@param buff
*/
func (bf *ByteBuffer) Append(buff []byte) {
	size := len(buff)
	if size == 0 {
		return
	}
	bf.write(buff, int32(size))
}

/*
WriteBytes
@Description: 写入n字节
@receiver bf
@param n
*/
func (bf *ByteBuffer) WriteBytes(n int32) {
	bf.writeIndex += n
}

/*
ReadBuff
@Description: 获取可读缓冲切片
@receiver bf
@return []byte
*/
func (bf *ByteBuffer) ReadBuff() []byte {
	buffLen := int32(len(bf.mBuffer))
	//如果读到了最后
	if bf.readIndex >= buffLen {
		return nil
	}
	//否则读取到最后
	return bf.mBuffer[bf.readIndex:]
}

/*
WriteBuff
@Description: 获取可写缓冲切片
@receiver bf
@return []byte
*/
func (bf *ByteBuffer) WriteBuff() []byte {
	buffLen := int32(len(bf.mBuffer))
	if bf.writeIndex >= buffLen {
		return nil
	}
	return bf.mBuffer[bf.writeIndex:]
}

/*
prependAbleBytes
@Description: 获取可读长度
@receiver bf
@return int32
*/
func (bf *ByteBuffer) prependAbleBytes() int32 {
	return bf.readIndex
}

/*
write
@Description: 写入数据，移动写指针
@receiver bf
@param buff: 写入的数据内容
@param len : 写入数据长度
*/
func (bf *ByteBuffer) write(buff []byte, len int32) {
	bf.EnsureWritableBytes(len)
	copy(bf.mBuffer[bf.writeIndex:], buff)
	bf.writeIndex = bf.writeIndex + len
}

/*
EnsureWritableBytes
@Description: 判断是能还能写len字节的数据
@receiver bf
@param len
*/
func (bf *ByteBuffer) EnsureWritableBytes(len int32) {
	if bf.writableBytes() < len {
		bf.grow(len)
	}
}

/*
grow
@Description: 扩容len字节长度
@receiver bf
@param len
*/
func (bf *ByteBuffer) grow(len int32) {
	//可写长度 + 已读长度（读过的长度和预置长度） < 指定长度 + 预置长度
	if bf.writableBytes()+bf.prependAbleBytes() < len+bf.reservedPrependSize {
		//当前可用长度 < len+预置长度
		//总容量*2 + 指定长度
		newCap := (bf.mCapacity << 1) + len
		buff := make([]byte, newCap)
		copy(buff, bf.mBuffer)
		bf.mCapacity = newCap
		bf.mBuffer = buff
	} else {
		//当前可用长度>=len+预置长度
		readable := bf.Length()
		//将未读部分放到新缓冲中
		copy(bf.mBuffer[bf.reservedPrependSize:], bf.mBuffer[bf.readIndex:bf.writeIndex])
		//读指针在预置位置
		bf.readIndex = bf.reservedPrependSize
		//写指针=读指针位置+未读长度
		bf.writeIndex = bf.readIndex + readable
	}
}

/*
NextBytes
@Description: 读取len字节数据，且移动读指针
@receiver bf
@param len
@return []byte
*/
func (bf *ByteBuffer) NextBytes(len int32) []byte {
	msgData := bf.mBuffer[bf.readIndex : bf.readIndex+len]
	bf.readIndex += len
	return msgData
}

/*
writableBytes
@Description: 获取可写长度
@receiver bf
@return int32
*/
func (bf *ByteBuffer) writableBytes() int32 {
	return bf.mCapacity - bf.writeIndex
}

/*
AppendInt64
@Description: 向缓冲区添加一个8字节的值
@receiver bf
@param x
*/
func (bf *ByteBuffer) AppendInt64(x int64) {
	buff := make([]byte, 8)
	if bf.littleEndian {
		binary.LittleEndian.PutUint64(buff, uint64(x))
	} else {
		binary.BigEndian.PutUint64(buff, uint64(x))
	}
	bf.write(buff, 8)
}

/*
AppendInt32
@Description: 向缓冲区添加一个4字节值
@receiver bf
@param x
*/
func (bf *ByteBuffer) AppendInt32(x int32) {
	buff := make([]byte, 4)
	if bf.littleEndian {
		binary.LittleEndian.PutUint32(buff, uint32(x))
	} else {
		binary.BigEndian.PutUint32(buff, uint32(x))
	}
	bf.write(buff, 4)
}

/*
AppendInt16
@Description: 向缓冲区添加一个2字节值
@receiver bf
@param x
*/
func (bf *ByteBuffer) AppendInt16(x int16) {
	buff := make([]byte, 2)
	if bf.littleEndian {
		binary.LittleEndian.PutUint16(buff, uint16(x))
	} else {
		binary.BigEndian.PutUint16(buff, uint16(x))
	}
	bf.write(buff, 2)
}

/*
ReadInt16
@Description: 一次读取2字节
@receiver bf
@return int16
*/
func (bf *ByteBuffer) ReadInt16() int16 {
	buff := bf.mBuffer[bf.readIndex : bf.readIndex+2]
	var result uint16
	if bf.littleEndian {
		result = binary.LittleEndian.Uint16(buff)
	} else {
		result = binary.BigEndian.Uint16(buff)
	}
	bf.Skip(2)
	return int16(result)
}

/*
ReadInt32
@Description: 读取4字节数据
@receiver bf
@return int32
*/
func (bf *ByteBuffer) ReadInt32() int32 {
	buff := bf.mBuffer[bf.readIndex : bf.readIndex+4]
	var result uint32
	if bf.littleEndian {
		result = binary.LittleEndian.Uint32(buff)
	} else {
		result = binary.BigEndian.Uint32(buff)
	}
	bf.Skip(4)
	return int32(result)
}

/*
ReadInt64
@Description: 读取8字节数据
@receiver bf
@return int64
*/
func (bf *ByteBuffer) ReadInt64() int64 {
	buff := bf.mBuffer[bf.readIndex : bf.readIndex+8]
	var result uint64
	if bf.littleEndian {
		result = binary.LittleEndian.Uint64(buff)
	} else {
		result = binary.BigEndian.Uint64(buff)
	}
	bf.Skip(8)
	return int64(result)
}

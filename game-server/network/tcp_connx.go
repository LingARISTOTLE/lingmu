package network

import (
	"encoding/binary"
	"fmt"
	spoor "github.com/LingARISTOTLE/lingolog"
	"io"
	"net"
	"reflect"
	"runtime/debug"
	"sync"
	"sync/atomic"
	"time"
)

type IConn interface {
	OnConnect()
	OnClose()
	OnMessage(*Message, *TcpConnX)
}

const timeoutTime = 30 // 连接通过验证的超时时间

/*
TcpConnX
@Description: 负责处理维护TCP连接，同时维护发送信道
*/
type TcpConnX struct {
	Conn        net.Conn         //连接
	Impl        IConn            //IConn的实现
	ConnID      int64            //连接id
	verify      int32            //验证状态
	closed      int32            //连接是否关闭
	stopped     chan bool        //是否停止连接
	signal      chan interface{} //发送队列
	lastSignal  chan interface{} //最终发送队列
	wgRW        sync.WaitGroup   //等待队列
	msgParser   *BufferPacker    //网络包缓冲处理类
	msgBuffSize int              //消息缓冲的大小
	logger      *spoor.Spoor     //日志
}

func NewTcpConnX(conn *net.TCPConn, msgBuffSize int, logger *spoor.Spoor) (*TcpConnX, error) {
	tcpConn := &TcpConnX{
		Conn:        conn,
		verify:      0,
		closed:      -1,
		stopped:     make(chan bool, 1),
		signal:      make(chan interface{}, 100),
		lastSignal:  make(chan interface{}, 1),
		wgRW:        sync.WaitGroup{},
		msgParser:   newInActionPacker(),
		msgBuffSize: msgBuffSize,
		logger:      logger,
	}
	//打开tcp的保活机制keepalive
	err := conn.SetKeepAlive(true)
	if err != nil {
		return nil, err
	}

	err = conn.SetKeepAlivePeriod(1 * time.Minute)
	if err != nil {
		return nil, err
	}

	//设置写缓冲
	err = conn.SetWriteBuffer(msgBuffSize)
	if err != nil {
		return nil, err
	}

	//设置读缓冲
	err = conn.SetReadBuffer(msgBuffSize)
	if err != nil {
		return nil, err
	}

	return tcpConn, nil
}

/*
Connect
@Description: 处理连接
@receiver c
*/
func (c *TcpConnX) Connect() {
	//cas将关闭状态由-1设置为0：表示运行
	if atomic.CompareAndSwapInt32(&c.closed, -1, 0) {
		c.wgRW.Add(1)
		go c.HandleRead()
		c.wgRW.Add(1)
		go c.HandleWrite()
	}
	timeout := time.NewTimer(time.Second * timeoutTime)

L:
	for {
		select {
		case <-timeout.C:
			if !c.Verified() {
				fmt.Printf("连接验证超时 ip addr %s", c.RemoteAddr())
				c.Close()
				break L
			}

		case <-c.stopped:
			break L
		}

	}
	timeout.Stop()
	c.wgRW.Wait()
	c.Impl.OnClose()
}

/*
HandleRead
@Description: 处理读操作
@receiver c
*/
func (c *TcpConnX) HandleRead() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("HandleRead panic ", err, "\n", string(debug.Stack()))
		}
	}()

	defer c.Close()

	//读协程不关闭，不能close
	defer c.wgRW.Done()

	for {
		data, err := c.msgParser.Read(c)
		if err != nil {
			if err != io.EOF {
				fmt.Printf("read message error: %v", err)
			}
		}
		message, err := c.msgParser.Unpack(data)
		c.Impl.OnMessage(message, c)
	}
}

/*
HandleWrite
@Description: 处理写数据
@receiver c
*/
func (c *TcpConnX) HandleWrite() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("HandleWrite panic", err, "\n", string(debug.Stack()))
		}
	}()

	defer c.Close()

	//服务关闭时，必须等待写协程执行到此处
	defer c.wgRW.Done()

	for {
		select {
		case signal := <-c.signal: // 普通消息
			data, ok := signal.([]byte)
			if !ok {
				fmt.Printf("write message %v error: msg is not bytes", reflect.TypeOf(signal))
				return
			}
			//写回
			err := c.msgParser.Write(c, data...)

			if err != nil {
				fmt.Printf("write message %v error: %v", reflect.TypeOf(signal), err)
				return
			}

		case signal := <-c.lastSignal: // 普通消息
			data, ok := signal.([]byte)
			if !ok {
				fmt.Printf("write message %v error: msg is not bytes", reflect.TypeOf(signal))
				return
			}

			err := c.msgParser.Write(c, data...)

			if err != nil {
				fmt.Printf("write message %v error: %v", reflect.TypeOf(signal), err)
				return
			}

			time.Sleep(2 * time.Second)
			return
		case <-c.stopped: //连接关闭通知
			return
		}

	}
}

/*
OnClose
@Description: 处理关闭连接方法
@receiver c
*/
func (c *TcpConnX) OnClose() {
	fmt.Printf("OnConnect 断开连接 local:%s remote:%s", c.LocalAddr(), c.RemoteAddr())
}

/*
Close
@Description: 关闭连接
@receiver c
*/
func (c *TcpConnX) Close() {
	if atomic.CompareAndSwapInt32(&c.closed, 0, 1) {
		//关闭连接
		c.Conn.Close()
		//关闭停止通道
		close(c.stopped)
	}
}

/*
OnConnect
@Description: 处理建立连接
@receiver c
*/
func (c *TcpConnX) OnConnect() {
	fmt.Printf("OnConnect 建立连接 local:%s remote:%s", c.LocalAddr(), c.RemoteAddr())
}

/*
Reset
@Description: 重制所有状态
@receiver c
*/
func (c *TcpConnX) Reset() {
	//cas获取关闭状态
	if atomic.LoadInt32(&c.closed) == -1 {
		return
	}
	//重制状态
	c.closed = -1
	c.verify = 0
	c.stopped = make(chan bool, 1)
	c.signal = make(chan interface{}, c.msgBuffSize)
	c.lastSignal = make(chan interface{}, 1)
	c.msgParser.reset()
}

/*
AsyncSendLastPacket
@Description: 缓存在最终发送队列里等待发送goroutine取出 (发送最后一个消息 发送会关闭tcp连接 终止tcp goroutine)
@receiver c
@param msgID
@param msg
@return bool
*/
func (c *TcpConnX) AsyncSendLastPacket(msgID uint16, msg interface{}) bool {
	data, err := c.msgParser.Pack(msgID, msg)
	if err != nil {
		fmt.Printf("AsyncSendLastPacket Pack msgID:%v and msg to bytes error:%v", msgID, err)
		return false
	}

	if uint32(len(data)) > c.msgParser.maxMsgLen {
		fmt.Printf("AsyncSendLastPacket 发送的消息包体过长 msgID:%v", msgID)
		return false
	}

	err = c.LastSignal(data)
	if err != nil {
		c.Close()
		fmt.Printf("%v", err)
		return false
	}

	return true

}

/*
AsyncSend
@Description: 异步发送消息
@receiver c
@param msgID
@param msg
@return bool
*/
func (c *TcpConnX) AsyncSend(msgID uint16, msg interface{}) bool {
	if c.IsShutdown() {
		return false
	}

	data, err := c.msgParser.Pack(msgID, msg)
	if err != nil {
		fmt.Printf("[AsyncSend] Pack msgID:%v and msg to bytes error:%v", msgID, err)
		return false
	}

	if uint32(len(data)) > c.msgParser.maxMsgLen {
		fmt.Printf("[AsyncSend] 发送的消息包体过长 msgID:%v", msgID)
		return false
	}

	//发送到发送通道
	err = c.Signal(data)
	if err != nil {
		c.Close()
		fmt.Printf("%v", err)
		return false
	}

	return true
}

/*
AsyncSendRowMsg
@Description: 按行发送（对最小长度没有限制）
@receiver c
@param data
@return bool
*/
func (c *TcpConnX) AsyncSendRowMsg(data []byte) bool {

	if c.IsShutdown() {
		return false
	}

	if uint32(len(data)) > c.msgParser.maxMsgLen {
		fmt.Printf("[AsyncSendRowMsg] 发送的消息包体过长 AsyncSendRowMsg")
		return false
	}

	err := c.Signal(data)
	if err != nil {
		c.Close()
		fmt.Printf("%v", err)
		return false
	}

	return true
}

/*
Signal
@Description: 将网络包发送到发送通道
@receiver c
@param signal
@return error
*/
func (c *TcpConnX) Signal(signal []byte) error {
	select {
	case c.signal <- signal:
		return nil
	default:
		cmd := binary.LittleEndian.Uint16(signal[2:4])
		//sprintf := fmt.Sprintf()
		return fmt.Errorf("Signal buffer full blocking connID:%v cmd:%v", c.ConnID, cmd)
	}
}

/*
LastSignal
@Description: 将网络包发送到最终发送通道
@receiver c
@param signal
@return error
*/
func (c *TcpConnX) LastSignal(signal []byte) error {
	select {
	case c.lastSignal <- signal:
		return nil
	default:
		cmd := binary.LittleEndian.Uint16(signal[2:4])
		//sprintf := fmt.Sprintf()
		return fmt.Errorf("LastSignal buffer full blocking connID:%v cmd:%v", c.ConnID, cmd)
	}
}

/*
RemoteAddr
@Description: 返回远程网络地址（如果已知）。
@receiver c
@return net.Addr
*/
func (c *TcpConnX) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

/*
LocalAddr
@Description: 返回本地网络地址（如果已知）
@receiver c
@return net.Addr
*/
func (c *TcpConnX) LocalAddr() net.Addr {
	return c.Conn.LocalAddr()
}

/*
Write
@Description: 将数据写回客户连接
@receiver c
@param b
@return int
@return error
*/
func (c *TcpConnX) Write(b []byte) (int, error) {
	return c.Conn.Write(b)
}

/*
Read
@Description: 将连接中的数据读取到字节数组中
@receiver c
@param b
@return int
@return error
*/
func (c *TcpConnX) Read(b []byte) (int, error) {
	return c.Conn.Read(b)
}

/*
IsShutdown
@Description: 返回关闭状态：1：Shutdown
@receiver c
@return bool
*/
func (c *TcpConnX) IsShutdown() bool {
	return atomic.LoadInt32(&c.closed) == 1
}

/*
IsClose
@Description: 只要关闭状态不是0，都是关闭
@receiver c
@return bool
*/
func (c *TcpConnX) IsClose() bool {
	return atomic.LoadInt32(&c.closed) != 0
}

/*
Verify
@Description: 验证连接：关闭-1，启动0，验证1
@receiver c
*/
func (c *TcpConnX) Verify() {
	atomic.CompareAndSwapInt32(&c.verify, 0, 1)
}

/*
Verified
@Description: 只要不是0，都是验证过了
@receiver c
@return bool
*/
func (c *TcpConnX) Verified() bool {
	return atomic.LoadInt32(&c.verify) != 0
}

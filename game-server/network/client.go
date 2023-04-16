package network

import (
	"fmt"
	spoor "github.com/LingARISTOTLE/lingolog"
	"net"
	"runtime/debug"
	"sync/atomic"
)

type Client struct {
	*TcpConnX
	Address         string                //连接地址
	ChMsg           chan *Message         //消息管道
	OnMessageCB     func(message *Packet) //客户端网络包处理（包含传输信息和连接）
	logger          *spoor.Spoor          //日志
	bufferSize      int                   //缓冲大小
	running         atomic.Value          //一个原子操作值，这里用来存储true和false来标记是否启动
	OnCloseCallBack func()                //关闭回调
	closed          int32                 //关闭状态
}

func NewClient(address string, connBuffSize int, logger *spoor.Spoor) *Client {
	client := &Client{
		bufferSize: connBuffSize,
		Address:    address,
		TcpConnX:   nil,
		logger:     logger,
	}
	client.running.Store(false)
	return client
}

/*
Dial
@Description: 获取TCP连接配置
@receiver c
@return *net.TCPConn
@return error
*/
func (c *Client) Dial() (*net.TCPConn, error) {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", c.Address)
	if err != nil {
		return nil, err
	}
	conn, err := net.DialTCP("tcp6", nil, tcpAddr)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

/*
Run
@Description: 启动客户端，启动协程专门用于获取连接
@receiver c
*/
func (c *Client) Run() {
	//获取连接配置
	conn, err := c.Dial()
	if err != nil {
		fmt.Printf("%v", err)
		return
	}

	tcpConnX, err := NewTcpConnX(conn, c.bufferSize, c.logger)
	if err != nil {
		fmt.Printf("%v", err)
		return
	}

	c.TcpConnX = tcpConnX
	c.Impl = c
	c.Reset()
	c.running.Store(true)
	go c.Connect()
}

/*
OnClose
@Description: 处理关闭连接
@receiver c
*/
func (c *Client) OnClose() {
	if c.OnCloseCallBack != nil {
		//调用关闭回调
		c.OnCloseCallBack()
	}
	//原子设置连接状态为false
	c.running.Store(false)
	//关闭连接
	c.TcpConnX.OnClose()
}

/*
OnMessage
@Description: 处理消息方法
@receiver c
@param data
@param conn
*/
func (c *Client) OnMessage(data *Message, conn *TcpConnX) {
	//验证连接
	c.Verify()

	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("[OnMessage] panic ", err, "\n", string(debug.Stack()))
		}
	}()

	//处理网络包
	c.OnMessageCB(&Packet{
		Msg:  data,
		Conn: conn,
	})
}

func (c *Client) Close() {
	if atomic.CompareAndSwapInt32(&c.closed, 0, 1) {
		c.Conn.Close()
		close(c.stopped)
	}
}

//func (c *Client) Write(conn net.Conn) {
//	//每秒生成一个chan Time到ticker中
//	//ticker := time.NewTicker(time.Second)
//
//	for {
//		select {
//		//每次发生事件：这里是测试，每秒读取一次定时器事件
//		//case <-ticker.C:
//		//	c.send(conn, &Message{
//		//		Id:   111,
//		//		Data: []byte("hello,world"),
//		//	})
//
//		case msg := <-c.ChMsg: //当控制台输入消息，那么发送
//			c.send(conn, msg)
//		}
//
//	}
//}
//
//func (c *Client) send(conn net.Conn, message *Message) {
//	err := conn.SetWriteDeadline(time.Now().Add(time.Second))
//	if err != nil {
//		//fmt.Println(err)
//		return
//	}
//	bytes, err := c.packer.Pack(message)
//	if err != nil {
//		//fmt.Println(err)
//		return
//	}
//
//	_, err = conn.Write(bytes)
//	if err != nil {
//		//fmt.Println(err)
//		return
//	}
//}
//
//func (c *Client) Read(conn net.Conn) {
//	for {
//		message, err := c.packer.UnPack(conn)
//
//		if _, ok := err.(net.Error); err != nil && ok {
//			fmt.Println(err)
//			continue
//		}
//
//		//调用消息方法，包装连接
//		c.OnMessage(&ClientPacket{
//			Msg:  message,
//			Conn: conn,
//		})
//
//		fmt.Println("resp message:", string(message.Data))
//
//	}
//}

//packer          IPacker                     //这里使用多态决定创建的对象

package network

import (
	"encoding/binary"
	"fmt"
	"net"
	"time"
)

type Client struct {
	Address string
	packer  *NormalPacker
	chMsg   chan *Message //消息管道
}

func NewClient(address string) *Client {
	return &Client{
		Address: address,
		packer: &NormalPacker{
			Order: binary.BigEndian, // 大端解读
		},
		chMsg: make(chan *Message, 1),
	}
}

func (c *Client) Run() {

	conn, err := net.Dial("tcp6", c.Address)
	if err != nil {
		fmt.Println(err)
		return
	}

	go c.Write(conn)
	go c.Read(conn)

}

func (c *Client) Write(conn net.Conn) {
	//每秒生成一个chan Time到ticker中
	ticker := time.NewTicker(time.Second)

	for {
		select {
		//每次发生事件：这里是测试，每秒读取一次定时器事件
		case <-ticker.C:
			c.send(conn, &Message{
				Id:   111,
				Data: []byte("hello,world"),
			})

		case msg := <-c.chMsg:
			c.send(conn, msg)
		}

	}
}

func (c *Client) send(conn net.Conn, message *Message) {
	err := conn.SetWriteDeadline(time.Now().Add(time.Second))
	if err != nil {
		fmt.Println(err)
		return
	}
	bytes, err := c.packer.Pack(message)
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = conn.Write(bytes)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func (c *Client) Read(conn net.Conn) {
	for {
		message, err := c.packer.UnPack(conn)

		if _, ok := err.(net.Error); err != nil && ok {
			fmt.Println(err)
			continue
		}

		fmt.Println("resp message:", string(message.Data))

	}
}

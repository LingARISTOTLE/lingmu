package main

import (
	"lingmu/game-server/network"
	"lingmu/game-server/network/protocol/gen/messageId"
)

type ClientManager struct {
	cli             *network.Client                        //客户端应用
	inputHandlers   map[string]InputHandler                //请求处理器集合
	messageHandlers map[messageId.MessageId]MessageHandler //消息处理器集合
	console         *ClientConsole                         //控制台
	chInput         chan *InputParam                       //客户端请求命令
}

/*
NewClient
@Description: 构造客户端对象，默认连接服务器的8023端口
@return *Client
*/
func NewClient() *ClientManager {
	c := &ClientManager{
		cli:             network.NewClient(":8023"),
		inputHandlers:   map[string]InputHandler{},
		messageHandlers: map[messageId.MessageId]MessageHandler{},
		console:         NewClientConsole(),
	}

	c.cli.OnMessage = c.OnMessage
	c.cli.ChMsg = make(chan *network.Message, 1) //消息通道缓存长度为1
	c.chInput = make(chan *InputParam, 1)        //客户端输入通道缓冲长度为1
	c.console.chInput = c.chInput
	return c
}

/*
Run
@Description: 启动客户端所有协程
@receiver c
*/
func (c *ClientManager) Run() {
	//启动协程，不断的读取控制台发送过来的信息，解析后交给客户端协程去发送
	go func() {
		for {
			select {
			case input := <-c.chInput:
				//fmt.Printf("cmd:%s,param:%v  <<<\t \n", input.Command, input.Param)
				inputHandler := c.inputHandlers[input.Command]
				if inputHandler != nil {
					inputHandler(input)
				}
				//发送消息
				c.cli.ChMsg <- &network.Message{
					Id:   uint64(messageId.MessageId_CSAddFriend),
					Data: []byte(input.Command),
				}
			}
		}
	}()
	//启动控制台协程，负责获取用户输入信息，并传递给包装协程
	go c.console.Run()
	//启动客户端服务，负责收发网络包
	go c.cli.Run()
}

func (c *ClientManager) OnMessage(packet *network.ClientPacket) {
	//把uint64转换为int32
	if handler, ok := c.messageHandlers[messageId.MessageId(packet.Msg.Id)]; ok {
		handler(packet)
	}
}

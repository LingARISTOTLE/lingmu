package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type ClientConsole struct {
	chInput chan *InputParam
}

/*
InputParam
@Description: 记录控制台输入信息，模拟客户端交互产生的网络包
*/
type InputParam struct {
	Command string   //指令，模拟网络包
	Param   []string //参数，模拟网络包参数
}

func NewClientConsole() *ClientConsole {
	c := &ClientConsole{}
	return c
}

func (c *ClientConsole) Run() {
	reader := bufio.NewReader(os.Stdin)
	for {
		//循环读取输入信息
		readString, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("input err,please check your input again!")
			continue
		}
		split := strings.Split(readString, " ")
		if len(split) == 0 {
			fmt.Println("input nil,please check your input again!")
			continue
		}

		//读取控制台输入信息
		in := &InputParam{
			Command: split[0],
			Param:   split[1:],
		}
		//将控制台输入信息发送到管道中
		c.chInput <- in
	}
}

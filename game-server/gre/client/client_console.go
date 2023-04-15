package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type ClientConsole struct {
	chInput chan *InputParam //用户输入命令
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

		strings.TrimSpace(readString)
		readString = strings.Replace(readString, "\n", "", -1)
		readString = strings.Replace(readString, "\r", "", -1)
		split := strings.Split(readString, " ")

		//读取控制台输入信息
		in := &InputParam{
			Command: split[0],
			Param:   split[1:],
		}
		//将控制台输入信息发送到管道中
		c.chInput <- in
	}
}

//split := strings.Split(readString, "|")
//newSlice := make([]string, 0)
//for i, s := range split {
//	//给每一段命令去除空格,\n,\r
//	split[i] = strings.TrimSpace(s)
//	split[i] = strings.Trim(s, "\n")
//	split[i] = strings.Trim(s, "\r")
//
//	if len(s) != 0 {
//		newSlice = append(newSlice, s)
//	}
//}
//
//if len(newSlice) == 0 {
//	fmt.Println("输入为nil，请检测输入!")
//	continue
//}

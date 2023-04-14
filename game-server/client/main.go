package main

import sugar "github.com/LingARISTOTLE/lingotools"

func main() {
	c := NewClient()
	//注册各种请求处理器
	c.InputHandlerRegister()
	c.MessageHandlerRegister()
	c.Run()
	sugar.WaitSignal(c.OnSystemSignal)
}

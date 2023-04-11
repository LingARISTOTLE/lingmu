package main

func main() {
	c := NewClient()
	//注册各种请求处理器
	c.InputHandlerRegister()
	c.Run()
	select {}
}

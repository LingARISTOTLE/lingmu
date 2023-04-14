package main

import (
	"fmt"
	sugar "github.com/LingARISTOTLE/lingotools"
	"lingmu/game-server/business/world"
)

func main() {

	//创建管理器组管理器
	world.MM = world.NewManagerHost()
	//启动组管理器
	world.MM.Run()
	fmt.Println("服务启动")

	sugar.WaitSignal(world.MM.OnSystemSignal)

}

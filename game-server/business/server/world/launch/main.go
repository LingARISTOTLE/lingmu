package main

import (
	"fmt"
	sugar "github.com/LingARISTOTLE/lingotools"
	"lingmu/game-server/business/server/world"
)

func main() {

	//创建管理器组管理器
	world.Oasis = world.NewWorld()
	//启动组管理器
	world.Oasis.Start()
	fmt.Println("服务启动")

	//主线线程阻塞
	sugar.WaitSignal(world.Oasis.OnSystemSignal)

}

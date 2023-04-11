package main

import "lingmu/game-server/world"

func main() {

	//创建管理器组管理器
	world.MM = world.NewManagerHost()
	//启动组管理器
	world.MM.Run()

}

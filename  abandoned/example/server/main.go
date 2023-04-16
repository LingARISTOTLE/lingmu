package main

import (
	"lingmu/game-server/network"
)

func main() {
	server := network.NewServer(":8023")
	go server.Run()
	select {}
}

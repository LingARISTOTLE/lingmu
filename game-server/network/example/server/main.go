package main

import (
	"lingmu/game-server/network"
)

func main() {
	server := network.NewServer(":8023")
	server.Run()
	select {}
}

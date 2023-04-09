package main

import "lingmu/game-server/network"

func main() {
	client := network.NewClient(":8023")
	client.Run()
	select {}
}

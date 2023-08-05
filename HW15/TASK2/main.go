package main

import (
	"HW15.2/observer"
	"time"
)

func main() {

	gameRoom := observer.NewItem("ROOM GAME & NETFLIX")

	observerFirst := &observer.Customer{"gamer123"}
	observerSecond := &observer.Customer{"noob_game_2015"}

	gameRoom.Register(observerFirst)
	gameRoom.Register(observerSecond)

	go gameRoom.SendTestEvents()
	//gameRoom.NotifyAll()
	go gameRoom.PoolNewPlayersInformation()

	time.Sleep(time.Second * 4)
}

package observer

import "fmt"

type Item struct {
	ObserverList []Observer
	NameGameRoom string
	PlayerEvents chan string
}

func NewItem(nameGameRoom string) *Item {
	return &Item{
		NameGameRoom: nameGameRoom,
		PlayerEvents: make(chan string),
	}
}

func (i *Item) PoolNewPlayersInformation() {
	for msg := range i.PlayerEvents {
		fmt.Println("********************")
		fmt.Printf("[INFO] New user %s joined game room\n", msg)
		fmt.Println("********************")
		for _, observer := range i.ObserverList {
			observer.Update(i.NameGameRoom, msg)
		}
	}
}

func (i *Item) SendTestEvents() {
	i.PlayerEvents <- "test_123_user_new"
	i.PlayerEvents <- "test_567_user_new"
	close(i.PlayerEvents)
}

func (i *Item) Register(o Observer) {
	i.ObserverList = append(i.ObserverList, o)
}

func (i *Item) Deregister(o Observer) {
	i.ObserverList = RemoveFromSlice(i.ObserverList, o)
}

func RemoveFromSlice(observerList []Observer, observerToRemove Observer) []Observer {
	observerListLength := len(observerList)
	for i, observer := range observerList {
		if observerToRemove.GetID() == observer.GetID() {
			observerList[observerListLength-1], observerList[i] = observerList[i], observerList[observerListLength-1]
			return observerList[:observerListLength-1]
		}
	}
	return observerList
}

func (i *Item) NotifyAll() {
	for _, observer := range i.ObserverList {
		observer.Update(i.NameGameRoom, "")
	}
}

func (i *Item) notifyAll() {
	for _, observer := range i.ObserverList {
		observer.Update(i.NameGameRoom, "")
	}
}

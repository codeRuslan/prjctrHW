package pubsub

import (
	"fmt"
	"sync"
)

type UserSubscribers map[string]*UserSubscriber

type Broker struct {
	userSubscribers UserSubscribers
	fileNames       map[string]UserSubscribers
	mut             sync.RWMutex
}

func NewBroker() *Broker {
	return &Broker{
		userSubscribers: UserSubscribers{},
		fileNames:       map[string]UserSubscribers{},
	}
}

func (b *Broker) AddUserSubscriber() *UserSubscriber {
	b.mut.Lock()
	defer b.mut.Unlock()

	id, s := CreateNewUserSubscriber()
	b.userSubscribers[id] = s
	return s
}

func (b *Broker) RemoverUserSubscriber(s *UserSubscriber) {
	for fileName := range s.fileNames {
		b.Unsubscribe(s, fileName)
	}

	b.mut.Lock()
	delete(b.userSubscribers, s.id)
	b.mut.Unlock()
}

func (b *Broker) Subscribe(us *UserSubscriber, fileName string) {
	b.mut.Lock()
	defer b.mut.Unlock()

	if b.fileNames[fileName] == nil {
		b.fileNames[fileName] = UserSubscribers{}
	}
	us.AddFileName(fileName)
	b.fileNames[fileName][us.id] = us
	fmt.Printf("%s Subscribed for fileName: %s\n", us.id, fileName)
}

func (b *Broker) Publish(fileName string, msg string) {
	b.mut.RLock()
	bFileNames := b.fileNames[fileName]
	b.mut.RUnlock()

	for _, s := range bFileNames {
		m := NewMessage(fileName, msg)

		if !s.active {
			return
		}

		go (func(us *UserSubscriber) {
			us.Signal(m)
		})(s)

	}
}

func (b *Broker) Unsubscribe(us *UserSubscriber, fileName string) {
	b.mut.RLock()
	defer b.mut.RUnlock()

	delete(b.fileNames[fileName], us.id)
	us.RemoveFileName(fileName)
	fmt.Printf("%s Unsubscribed for fileName: %s\n", us.id, fileName)
}

func (b *Broker) GetUserSubscribers(fileName string) int {
	b.mut.RLock()
	defer b.mut.RUnlock()
	return len(b.fileNames[fileName])
}

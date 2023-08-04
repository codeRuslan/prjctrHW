package pubsub

import (
	"fmt"
	"math/rand"
	"sync"
)

type UserSubscriber struct {
	id        string
	messages  chan *Message
	fileNames map[string]bool
	active    bool
	mutex     sync.RWMutex
}

func CreateNewUserSubscriber() (string, *UserSubscriber) {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	id := make([]byte, 6)
	for i := range id {
		id[i] = letterBytes[rand.Intn(len(letterBytes))]
	}

	return string(id), &UserSubscriber{
		id:        string(id),
		messages:  make(chan *Message),
		fileNames: map[string]bool{},
		active:    true,
	}
}

func (us *UserSubscriber) AddFileName(fileName string) {
	us.mutex.RLock()
	defer us.mutex.RUnlock()
	us.fileNames[fileName] = true
}

func (us *UserSubscriber) RemoveFileName(fileName string) {
	us.mutex.RLock()
	defer us.mutex.RUnlock()
	delete(us.fileNames, fileName)
}

func (us *UserSubscriber) GetFileNames() []string {
	us.mutex.RLock()
	defer us.mutex.RUnlock()
	fileNames := []string{}

	for k, _ := range us.fileNames {
		fileNames = append(fileNames, k)
	}
	return fileNames
}

func (us *UserSubscriber) Destroy() {
	us.mutex.RLock()
	defer us.mutex.RUnlock()

	us.active = false
	close(us.messages)
}

func (us *UserSubscriber) Signal(msg *Message) {
	us.mutex.RLock()
	defer us.mutex.RUnlock()
	if us.active {
		us.messages <- msg
	}
}

func (us *UserSubscriber) Listen() {
	for {
		if msg, ok := <-us.messages; ok {
			fmt.Printf("Filename: %s has been changed the following way: %s \n", msg.fileName, msg.action) // msg.fileName
		}
	}
}

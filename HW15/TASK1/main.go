package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"log"
	"os"
	"path/filepath"
	"testPlayground2/pubsub"
	"time"
)

func createAvailableFileNames(path string) map[string]string {
	availableFileNames := map[string]string{}
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			fileName := filepath.Base(path)
			availableFileNames[fileName] = path
		}
		return nil
	})

	if err != nil {
		fmt.Println("Error while walking through the directory:", err)
		return nil
	}

	// Print the map
	for fileName, filePath := range availableFileNames {
		fmt.Printf("%s: %s\n", fileName, filePath)
	}
	return availableFileNames
}

var ListFiles map[string]string

func main() {
	ListFiles = createAvailableFileNames("/Users/ruslanpilipyuk/GolandProjects/testPlayground2/playgroundfolder/")
	Broker := pubsub.NewBroker()

	s1 := Broker.AddUserSubscriber() // first user
	for k, _ := range ListFiles {
		Broker.Subscribe(s1, ListFiles[k])
	}

	s2 := Broker.AddUserSubscriber()
	for k, _ := range ListFiles {
		Broker.Subscribe(s2, ListFiles[k])
	}

	go (func() {
		time.Sleep(3 * time.Second)
		fmt.Printf("Total subscribers for topic test1 file is %v\n", Broker.GetUserSubscribers(ListFiles["test1"]))
	})()

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal("NewWatcher Failed: ", err.Error())
	}
	defer watcher.Close()

	done := make(chan bool)

	go func() {
		defer close(done)

		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				for _, v := range ListFiles {
					if event.Name == v {
						Broker.Publish(v, string(event.Op))
					}
				}
				//log.Printf("%s %s\n", event.Name, event.Op)
			}
		}
	}()

	err = watcher.Add("/Users/ruslanpilipyuk/GolandProjects/testPlayground2/playgroundfolder/")

	if err != nil {
		log.Fatal("Add failed:", err)
	}

	go s1.Listen()
	go s2.Listen()

	<-done

	fmt.Scanln()
	fmt.Println("Done!")
}

/*
func FileWatcher() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal("NewWatcher Failed: ", err.Error())
	}
	defer watcher.Close()

	done := make(chan bool)

	go func() {
		defer close(done)

		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				for _, v := range ListFiles {
					if event.Name == v {
						log.Printf("FOUND EQUALITY: %s %s\n", event.Name, event.Op)

						go

					}
				}
				log.Printf("%s %s\n", event.Name, event.Op)
			}
		}
	}()

	err = watcher.Add("/Users/ruslanpilipyuk/GolandProjects/testPlayground2/playgroundfolder/")

	if err != nil {
		log.Fatal("Add failed:", err)
	}

	<-done

}

*/

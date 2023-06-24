package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	numChannel := make(chan int)
	responseChannel := make(chan int)

	go func() {
		for i := 0; i < 3; i++ {
			numChannel <- rand.Intn(10)
		}
		close(numChannel)
	}()

	go func() {
		var sum int
		for num := range numChannel {
			fmt.Println(num)
			sum += num
		}
		responseChannel <- sum / 3
		close(responseChannel)
	}()

	go func() {
		answer := <-responseChannel
		fmt.Println("Answer from response channel:", answer)
	}()

	time.Sleep(time.Second)
}

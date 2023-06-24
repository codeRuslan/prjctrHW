package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	amountOfNumsGeneated := 3
	firstChannel := make(chan int)
	numSlice := make([]int, 0, amountOfNumsGeneated) // Specify the length and capacity of the slice
	min := 1
	max := 100

	go func() {
		for i := 0; i < amountOfNumsGeneated; i++ {
			firstChannel <- rand.Intn(max-min+1) + min
		}
		msgMinSlice := <-firstChannel
		msgMaxSlice := <-firstChannel
		fmt.Printf("Max number: %d. Min number: %d", msgMaxSlice, msgMinSlice)
	}()

	go func() {
		for i := 0; i < amountOfNumsGeneated; i++ {
			num := <-firstChannel
			numSlice = append(numSlice, num)
		}

		minSlice := numSlice[0]
		maxSlice := numSlice[0]
		fmt.Println("------")
		fmt.Println("How list looks like now")
		for i, v := range numSlice {
			fmt.Println(i, v)
		}
		fmt.Println("------")

		for i := 0; i < len(numSlice); i++ {
			if numSlice[i] < minSlice {
				minSlice = numSlice[i]
			}
			if numSlice[i] > maxSlice {
				maxSlice = numSlice[i]
			}
		}

		firstChannel <- minSlice
		firstChannel <- maxSlice
		close(firstChannel)
	}()

	time.Sleep(time.Second)
}

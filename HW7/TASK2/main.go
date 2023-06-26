package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	amountOfNumsGenerated := rand.Intn(10) + 1 // Генеруємо випадкове число від 1 до 10
	firstChannel := make(chan int)

	go func() {
		min := 1
		max := 100

		for i := 0; i < amountOfNumsGenerated; i++ {
			firstChannel <- rand.Intn(max-min+1) + min
		}

		msgMinSlice := <-firstChannel
		msgMaxSlice := <-firstChannel
		fmt.Printf("Max number: %d. Min number: %d\n", msgMaxSlice, msgMinSlice)
	}()

	go func() {
		numSlice := make([]int, 0, amountOfNumsGenerated) // Вказуємо довжину та ємність зрізу

		for i := 0; i < amountOfNumsGenerated; i++ {
			num := <-firstChannel
			numSlice = append(numSlice, num)
		}

		minSlice := numSlice[0]
		maxSlice := numSlice[0]
		fmt.Println("------")
		fmt.Println("How the list looks now:")
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

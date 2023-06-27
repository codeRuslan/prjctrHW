package main

import (
	"context"
	"flag"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Info struct {
	Name         string
	GoodName     string
	AmountOfGood int
}

func main() {
	ctx := context.Background()
	ctx, stop := context.WithCancel(ctx)
	channelInfo := make(chan Info, 100) // Buffered channel for better performance
	count := flag.Int("amount", 5, "Amount of customer info to process")
	flag.Parse()

	names := []string{"Luiz", "Andrew", "Jame", "Anastasia", "John", "Margo", "Krio", "Maria", "Philip", "Julia"}
	goodNames := []string{"ball", "chair", "table", "knife", "tires", "tent", "shelf", "water"}
	prices := []int{10, 49, 89, 15, 29, 99, 79, 1}

	goodPrices := make(map[string]int)
	for i, name := range goodNames {
		goodPrices[name] = prices[i]
	}

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		rand.Seed(time.Now().UnixNano())

		for i := 0; i < *count; i++ {
			randIndex := rand.Intn(len(names))
			randName := names[randIndex]

			randIndex = rand.Intn(len(goodNames))
			randGood := goodNames[randIndex]

			randAmountOfGood := rand.Intn(10)

			GeneratedInfo := Info{
				Name:         randName,
				GoodName:     randGood,
				AmountOfGood: randAmountOfGood,
			}
			channelInfo <- GeneratedInfo
		}

		stop()
	}()

	go func() {
		defer wg.Done()
		sumOfOrders := 0
		for {
			select {
			case clientInfo := <-channelInfo:
				sumOfOrders += goodPrices[clientInfo.GoodName] * clientInfo.AmountOfGood
				fmt.Println(clientInfo)
			case <-ctx.Done():
				fmt.Println("Sum of all goods:")
				fmt.Println(sumOfOrders)
				return
			}
		}
	}()

	wg.Wait()
}

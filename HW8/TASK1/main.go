package main

import (
	"context"
	"flag"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Order struct {
	Name         string
	GoodName     string
	AmountOfGood int
}

type PriceCalculator struct {
	Prices map[string]int
}

func (pc *PriceCalculator) CalculateTotalPrice(order Order) int {
	return pc.Prices[order.GoodName] * order.AmountOfGood
}

func main() {
	ctx := context.Background()
	ctx, stop := context.WithCancel(ctx)
	channelInfo := make(chan Order, 100) // Buffered channel for better performance
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

			GeneratedOrder := Order{
				Name:         randName,
				GoodName:     randGood,
				AmountOfGood: randAmountOfGood,
			}
			channelInfo <- GeneratedOrder
		}
		time.Sleep(time.Second)
		stop()
	}()

	go func() {
		defer wg.Done()
		pc := PriceCalculator{Prices: goodPrices}
		sumOfOrders := 0

		for {
			select {
			case order := <-channelInfo:
				orderPrice := pc.CalculateTotalPrice(order)
				sumOfOrders += orderPrice
				fmt.Println(order)
			case <-ctx.Done():
				fmt.Println("Sum of all goods:")
				fmt.Println(sumOfOrders)
				return
			}
		}
	}()

	wg.Wait()

}

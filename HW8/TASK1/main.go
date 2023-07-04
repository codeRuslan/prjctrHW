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

type OrderGenerator struct {
	Names     []string
	GoodNames []string
}

func (og *OrderGenerator) GenerateOrder() Order {
	rand.Seed(time.Now().UnixNano())

	randIndex := rand.Intn(len(og.Names))
	randName := og.Names[randIndex]

	randIndex = rand.Intn(len(og.GoodNames))
	randGood := og.GoodNames[randIndex]

	randAmountOfGood := rand.Intn(10)

	return Order{
		Name:         randName,
		GoodName:     randGood,
		AmountOfGood: randAmountOfGood,
	}
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

	orderGenerator := OrderGenerator{
		Names:     []string{"Luiz", "Andrew", "Jame", "Anastasia", "John", "Margo", "Krio", "Maria", "Philip", "Julia"},
		GoodNames: []string{"ball", "chair", "table", "knife", "tires", "tent", "shelf", "water"},
	}

	goodPrices := map[string]int{
		"ball":  10,
		"chair": 49,
		"table": 89,
		"knife": 15,
		"tires": 29,
		"tent":  99,
		"shelf": 79,
		"water": 1,
	}

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()

		for i := 0; i < *count; i++ {
			GeneratedOrder := orderGenerator.GenerateOrder()
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

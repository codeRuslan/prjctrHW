package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Створити програму для симуляції групи людей, які одночасно грають в ігри на великому екрані.
//Програма має використовувати горутину-генератор, який кожні 10 секунд генерує новий ігровий
//раунд та відправляє його до горутин-гравців через канал. Гравці отримують новий ігровий раунд та
//вводять свої відповіді через окремий канал. Далі горутина-лічильник перевіряє правильність
//відповідей та повертає результат у головну горутину через окремий канал. Якщо в програмі
//виникає помилка або користувач перериває програму, то програма має коректно завершувати роботу з використанням контексту.

func main() {
	iterationCount := 1
	playerCount := 2
	questionsChannel := make(chan int)
	calculatorChannel := make(chan int)
	outputChannel := make(chan int)
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	rand.Seed(time.Now().UnixNano())
	ticker := time.NewTicker(10 * time.Second)

	go func(ctxFunc context.Context) { // generator
		for {
			select {
			case <-ticker.C:
				generatedAnswer := rand.Intn(10)
				for i := 0; i < playerCount+1; i++ {
					questionsChannel <- generatedAnswer
					//fmt.Println(generatedAnswer)
				}
			case <-ctxFunc.Done():
				fmt.Println("[Generator] Received Cancel() shutting down...")
				return
			}
		}
	}(ctx)

	go func(ctxFunc context.Context) { // Player1
		for {
			select {
			case <-ticker.C:
				answerFromPlayerChannel := <-questionsChannel
				fmt.Printf("[ITERATION #%d][Player1] Got response: %d\n", iterationCount, answerFromPlayerChannel)
				calculatorChannel <- answerFromPlayerChannel // Player sending out right answer
			case <-ctxFunc.Done():
				fmt.Println("[Player1] Received Cancel() shutting down...")
				return
			}
		}
	}(ctx)

	go func(ctxFunc context.Context) { // Player2
		for {
			select {
			case <-ticker.C:
				answerFromPlayerChannel := <-questionsChannel
				fmt.Printf("[ITERATION #%d][Player2] Got response: %d\n", iterationCount, answerFromPlayerChannel)
				calculatorChannel <- answerFromPlayerChannel // Player sending out right answer
			case <-ctxFunc.Done():
				fmt.Println("[Player2] Received Cancel() shutting down...")
				return
			}
		}
	}(ctx)

	go func(ctxFunc context.Context) { // Calculator Of Right Answers
		countValidAnswersFromPlayers := 0
		for {
			select {
			case <-ticker.C:
				answerFromPlayerChannel := <-questionsChannel
				for i := 0; i < 2; i++ {
					guessFromPlayer := <-calculatorChannel
					if guessFromPlayer == answerFromPlayerChannel {
						countValidAnswersFromPlayers++
					}
				}
				outputChannel <- countValidAnswersFromPlayers
			case <-ctxFunc.Done():
				fmt.Println("[Calculator] Received Cancel() shutting down...")
				return
			}
		}
	}(ctx)

	go func() { // This will trigger cancel of ctx, it works by receiving termination signal and starting shutdown process
		<-signalChan
		cancel()
		time.Sleep(time.Second * 3)
		fmt.Println("[Information] Sucessfully managed to gracefuly shut down all components")
		fmt.Println("Exiting....")
		os.Exit(1)
	}()

	for {
		fmt.Printf("[ITERATION #%d] Count of valid answers equals to => %d\n", iterationCount, <-outputChannel)
		iterationCount++
		fmt.Println("----------------- NEXT ITERATION HAS STARTED -----------------")
	}

}

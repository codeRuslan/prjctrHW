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

	go generateQuestions(ctx, ticker, questionsChannel, playerCount+1)
	go processPlayer(ctx, ticker, questionsChannel, calculatorChannel, iterationCount, "Player1")
	go processPlayer(ctx, ticker, questionsChannel, calculatorChannel, iterationCount, "Player2")
	go calculateAnswers(ctx, ticker, questionsChannel, calculatorChannel, outputChannel)

	go func() {
		<-signalChan
		cancel()
		time.Sleep(time.Second * 3)
		fmt.Println("[Information] Successfully managed to gracefully shut down all components")
		fmt.Println("Exiting....")
		os.Exit(1)
	}()

	for {
		fmt.Printf("[ITERATION #%d] Count of valid answers equals to => %d\n", iterationCount, <-outputChannel)
		iterationCount++
		fmt.Println("----------------- NEXT ITERATION HAS STARTED -----------------")
	}
}

func generateQuestions(ctx context.Context, ticker *time.Ticker, questionsChannel chan<- int, playerCount int) {
	for {
		select {
		case <-ticker.C:
			generatedAnswer := rand.Intn(10)
			for i := 0; i < playerCount; i++ {
				questionsChannel <- generatedAnswer
			}
		case <-ctx.Done():
			fmt.Println("[Generator] Received Cancel() shutting down...")
			return
		}
	}
}

func processPlayer(ctx context.Context, ticker *time.Ticker, questionsChannel <-chan int, calculatorChannel chan<- int, iterationCount int, playerName string) {
	for {
		select {
		case <-ticker.C:
			answerFromPlayerChannel := <-questionsChannel
			fmt.Printf("[ITERATION #%d][%s] Got response: %d\n", iterationCount, playerName, answerFromPlayerChannel)
			calculatorChannel <- answerFromPlayerChannel
		case <-ctx.Done():
			fmt.Printf("[%s] Received Cancel() shutting down...\n", playerName)
			return
		}
	}
}

func calculateAnswers(ctx context.Context, ticker *time.Ticker, questionsChannel <-chan int, calculatorChannel <-chan int, outputChannel chan<- int) {
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
		case <-ctx.Done():
			fmt.Println("[Calculator] Received Cancel() shutting down...")
			return
		}
	}
}

package main

import (
	"fmt"
	"strconv"
)

var rows = []string{"A", "B", "C"}
var columns = []int{1, 2, 3}
var gameWon = false
var lastTurn string
var turnCounter = 0

func GetPlayingBoard(playingBoard map[string]string) {
	fmt.Println("  ===================")
	fmt.Println("     1 ---- 2 --- 3")
	fmt.Println("  -------------------")
	for _, row := range rows {
		fmt.Print(row)
		for i := 1; i <= 3; i++ {
			if i == 1 {
				fmt.Print("  |")
			}
			indexRow := playingBoard[row+(strconv.Itoa(i))]
			fmt.Print("  " + indexRow + "  |")
			if i == 3 {
				fmt.Println("")
			}
		}
		fmt.Println("  -------------------")
	}
	fmt.Println("===================")
}

func PutXO(coordinate string, playingBoard map[string]string, symbol string) {
	playingBoard[coordinate] = symbol
}

func PutPlayerTurn(playerType string, playingBoard map[string]string) {
	turnCounter++
	var coordinate string
	fmt.Println("\nPlayer " + playerType + " your turn:")
	fmt.Println("------------")
	for {
		fmt.Println("Choose coordinates to place " + playerType + ":")
		fmt.Scanln(&coordinate)
		if !isValidCoordinate(playingBoard, coordinate) {
			continue
		}
		PutXO(coordinate, playingBoard, playerType)
		fmt.Println("State of Playingboard after Player " + playerType + " did their turn:")
		fmt.Println("Place " + playerType + " on coordinate " + coordinate)
		GetPlayingBoard(playingBoard)
		break
	}
	lastTurn = playerType
	CheckWinCondition(playingBoard, coordinate)
}

func isValidCoordinate(playingBoard map[string]string, coordinate string) bool {
	_, ok := playingBoard[coordinate]
	if !ok {
		fmt.Println("There is no such coordinate!")
		fmt.Println("Try again")
		return false
	}

	if playingBoard[coordinate] != "-" {
		fmt.Println("You cannot use that coordinate, as it is already occupied!")
		fmt.Println("Try again")
		return false
	}
	return true
}

func CheckWinCondition(playingBoard map[string]string, coordinate string) {
	// CHECK horizontal win
	for _, row := range rows {
		if row == string(coordinate[0]) {
			if playingBoard[row+"1"] == playingBoard[row+"2"] && playingBoard[row+"2"] == playingBoard[row+"3"] {
				gameWon = true
			}
		}
	}

	// CHECK vertical win
	keyColumnToInt, _ := strconv.Atoi(string(coordinate[1]))
	for _, column := range columns {
		if column == keyColumnToInt {
			if playingBoard["A"+strconv.Itoa(column)] == playingBoard["B"+strconv.Itoa(column)] && playingBoard["B"+strconv.Itoa(column)] == playingBoard["C"+strconv.Itoa(column)] {
				gameWon = true
			}
		}
	}

	// CHECK Diagonal win
	if playingBoard["A1"] != "-" && playingBoard["B2"] != "-" && playingBoard["C3"] != "-" {
		if playingBoard["A1"] == playingBoard["B2"] && playingBoard["B2"] == playingBoard["C3"] {
			gameWon = true
		}
	}
	if playingBoard["A3"] != "-" && playingBoard["B2"] != "-" && playingBoard["C1"] != "-" {
		if playingBoard["A3"] == playingBoard["B2"] && playingBoard["B2"] == playingBoard["C1"] {
			gameWon = true
		}
	}
}

func main() {
	playingBoard := make(map[string]string)
	for _, row := range rows {
		for _, column := range columns {
			columnStr := strconv.Itoa(column)
			playingBoard[row+columnStr] = "-"
		}
	}

	fmt.Println("Game has been initiated:\n")
	GetPlayingBoard(playingBoard)

	for !gameWon && turnCounter < 9 {
		PutPlayerTurn("X", playingBoard)
		if turnCounter == 9 {
			break
		}
		PutPlayerTurn("O", playingBoard)
	}

	fmt.Println("Game has been finished!")
	if gameWon {
		fmt.Println(lastTurn + " Player has won the game")
	} else {
		fmt.Println("The game ended in a draw")
	}
}

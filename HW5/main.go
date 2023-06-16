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
		// var coordinate string
		fmt.Scanln(&coordinate)
		if playingBoard[coordinate] != "-" {
			fmt.Println("You cannot use that coordinate, as it is busy already!")
			fmt.Println("Try one more time")
		} else if _, ok := playingBoard[coordinate]; !ok {
			fmt.Println("There are no such coordinate!")
			fmt.Println("Try one more time")
		} else {
			PutXO(coordinate, playingBoard, playerType)
			fmt.Println("State of Playingboard after Player " + playerType + " did his turn:")
			fmt.Println("Place " + playerType + " on coordinate " + coordinate)
			GetPlayingBoard(playingBoard)
			break
		}
	}
	lastTurn = playerType
	CheckWinCondition(playingBoard, coordinate)
}

func CheckWinCondition(playingBoard map[string]string, coordinate string) {
	// CHECK horizontal win
	for _, row := range rows {
		if row == string(coordinate[0]) {
			if (playingBoard[row+"1"] == playingBoard[row+"2"]) && (playingBoard[row+"2"] == playingBoard[row+"3"]) {
				gameWon = true
			}
		}
	}

	// CHECK vertical win

	KeyColumnToInt, _ := strconv.Atoi(string(coordinate[1]))
	for _, column := range columns {
		if column == KeyColumnToInt {
			if (playingBoard["A"+strconv.Itoa(column)] == playingBoard["B"+strconv.Itoa(column)]) && (playingBoard["B"+strconv.Itoa(column)] == playingBoard["C"+strconv.Itoa(column)]) {
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
	var playingBoard = make(map[string]string)
	for _, row := range rows {
		for _, column := range columns {
			columnStr := strconv.Itoa(column)
			playingBoard[row+columnStr] = "-"
		}
	}

	for {
		if gameWon == false {
			fmt.Println("Game has been initiated:\n")
			GetPlayingBoard(playingBoard)

			PutPlayerTurn("X", playingBoard)
			if turnCounter == 9 {
				fmt.Println("Game has been finished!")
				fmt.Println("You got a draw")
				break
			}

			if gameWon != false {
				fmt.Println("Game has been finished!")
				fmt.Println(lastTurn + " Player has won a game")
				break
			}
			PutPlayerTurn("O", playingBoard)

		} else {
			fmt.Println("Game has been finished!")
			fmt.Println(lastTurn + " Player has won a game")
			break
		}
	}

}

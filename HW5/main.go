package main

import (
	"fmt"
	"strconv"
)

var rows = []string{"A", "B", "C"}
var columns = []int{1, 2, 3}

type Game struct {
	playingBoard map[string]string
	gameWon      bool
	lastTurn     string
	turnCounter  int
}

func NewGame() *Game {
	playingBoard := make(map[string]string)
	for _, row := range rows {
		for _, column := range columns {
			columnStr := strconv.Itoa(column)
			playingBoard[row+columnStr] = "-"
		}
	}

	return &Game{
		playingBoard: playingBoard,
		gameWon:      false,
		lastTurn:     "",
		turnCounter:  0,
	}
}

func (g *Game) GetPlayingBoard() {
	fmt.Println("  ===================")
	fmt.Println("     1 ---- 2 --- 3")
	fmt.Println("  -------------------")
	for _, row := range rows {
		fmt.Print(row)
		for i := 1; i <= 3; i++ {
			if i == 1 {
				fmt.Print("  |")
			}
			indexRow := g.playingBoard[row+(strconv.Itoa(i))]
			fmt.Print("  " + indexRow + "  |")
			if i == 3 {
				fmt.Println("")
			}
		}
		fmt.Println("  -------------------")
	}
	fmt.Println("===================")
}

func (g *Game) PutXO(coordinate string, symbol string) {
	g.playingBoard[coordinate] = symbol
}

func (g *Game) PutPlayerTurn(playerType string) {
	g.turnCounter++
	var coordinate string
	fmt.Println("\nPlayer " + playerType + " your turn:")
	fmt.Println("------------")
	for {
		fmt.Println("Choose coordinates to place " + playerType + ":")
		fmt.Scanln(&coordinate)
		if !g.isValidCoordinate(coordinate) {
			continue
		}
		g.PutXO(coordinate, playerType)
		fmt.Println("State of Playingboard after Player " + playerType + " did their turn:")
		fmt.Println("Place " + playerType + " on coordinate " + coordinate)
		g.GetPlayingBoard()
		break
	}
	g.lastTurn = playerType
	g.CheckWinCondition(coordinate)
}

func (g *Game) isValidCoordinate(coordinate string) bool {
	_, ok := g.playingBoard[coordinate]
	if !ok {
		fmt.Println("There is no such coordinate!")
		fmt.Println("Try again")
		return false
	}

	if g.playingBoard[coordinate] != "-" {
		fmt.Println("You cannot use that coordinate, as it is already occupied!")
		fmt.Println("Try again")
		return false
	}
	return true
}

func (g *Game) CheckWinCondition(coordinate string) {
	// CHECK horizontal win
	for _, row := range rows {
		if row == string(coordinate[0]) {
			if g.playingBoard[row+"1"] == g.playingBoard[row+"2"] && g.playingBoard[row+"2"] == g.playingBoard[row+"3"] && g.playingBoard[row+"1"] != "-" {
				g.gameWon = true
			}
		}
	}

	// CHECK vertical win
	keyColumnToInt, _ := strconv.Atoi(string(coordinate[1]))
	for _, column := range columns {
		if column == keyColumnToInt {
			if g.playingBoard["A"+strconv.Itoa(column)] == g.playingBoard["B"+strconv.Itoa(column)] && g.playingBoard["B"+strconv.Itoa(column)] == g.playingBoard["C"+strconv.Itoa(column)] && g.playingBoard["A"+strconv.Itoa(column)] != "-" {
				g.gameWon = true
			}
		}
	}

	// CHECK Diagonal win
	if g.playingBoard["A1"] != "-" && g.playingBoard["B2"] != "-" && g.playingBoard["C3"] != "-" {
		if g.playingBoard["A1"] == g.playingBoard["B2"] && g.playingBoard["B2"] == g.playingBoard["C3"] {
			g.gameWon = true
		}
	}
	if g.playingBoard["A3"] != "-" && g.playingBoard["B2"] != "-" && g.playingBoard["C1"] != "-" {
		if g.playingBoard["A3"] == g.playingBoard["B2"] && g.playingBoard["B2"] == g.playingBoard["C1"] {
			g.gameWon = true
		}
	}
}

func (g *Game) Play() {
	fmt.Println("Game has been initiated:\n")
	g.GetPlayingBoard()

	for g.IsGameFinished() {
		g.PutPlayerTurn("X")
		if g.turnCounter == 9 || g.gameWon {
			break
		}
		g.PutPlayerTurn("O")
	}

	fmt.Println("Game has been finished!")
	if g.gameWon {
		fmt.Println(g.lastTurn + " Player has won the game")
	} else {
		fmt.Println("The game ended in a draw")
	}
}

func (g *Game) IsGameFinished() bool {
	return !g.gameWon && g.turnCounter < 9
}

func main() {
	game := NewGame()
	game.Play()
}

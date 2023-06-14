package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var inputStringForSearch string

type Notebook struct {
	text [][]string
}

func (n *Notebook) PutTextNotebook(inputText [][]string) {
	n.text = inputText
}

func (n *Notebook) GetTextNotebook() [][]string {
	return n.text
}

func (n *Notebook) SearchForNotebookText(inputText string) {
	for i, row := range n.text {
		rowText := strings.Join(row, " ")
		if strings.Contains(rowText, inputText) {
			fmt.Println("Row number for your search query is --> ", i+1)
			fmt.Println("Row content: ", row)
		}
	}

	if len(n.text) == 0 {
		fmt.Println("The notebook is empty. Please add some text first.")
	} else {
		fmt.Println("Search complete.")
	}
}

func (n *Notebook) SearchForNotebookRow(inputRow []string) {
	foundRow := false
	checkRowsEqual := true

	for i, row := range n.text {
		for b, element := range inputRow {
			if element != row[b] {
				checkRowsEqual = false
			}
		}
		if len(row) == len(inputRow) && checkRowsEqual {
			foundRow = true
			fmt.Println("Row number for your search query is --> ", i+1)
		}
	}

	if !foundRow {
		fmt.Println("We could not find the row you have been searching for. Try again!")
	}
}

func (n *Notebook) AddRowToNotebook(row []string) {
	n.text = append(n.text, row)
}

func main() {
	noteBookText := [][]string{}
	scanner := bufio.NewScanner(os.Stdin)

	mainNotebook := Notebook{
		noteBookText,
	}

	fmt.Println("Text inside Notebook:")
	fmt.Println("******************************")
	fmt.Println(mainNotebook.GetTextNotebook())
	fmt.Println("******************************")

	fmt.Println("Here are the results of the search for a row inside Notebook Text:")
	testInput := []string{"Hello", "World"} // To test if SearchForNotebookRow works fine
	mainNotebook.SearchForNotebookRow(testInput)

	fmt.Println("Input word combinations you are looking for (press Enter to exit):")
	for scanner.Scan() {
		text := scanner.Text()

		if text == "" {
			break
		}

		sliceText := strings.Split(text, " ")
		mainNotebook.AddRowToNotebook(sliceText)
	}

	fmt.Println("Text inside Notebook after adding new rows:")
	fmt.Println("******************************")
	fmt.Println(mainNotebook.GetTextNotebook())
	fmt.Println("******************************")

	fmt.Println("Input a word combination you are looking for:")
	fmt.Scan(&inputStringForSearch)
	mainNotebook.SearchForNotebookText(inputStringForSearch)
}

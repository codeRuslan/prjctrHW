package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strings"
)

/*
Завдання про текстовий редактор. Створити slice string з текстом, який користувач вводить у текстовий редактор.
Написати функцію, яка приймає на вхід рядок та знаходить у текстовому редакторі всі рядки, які містять цей рядок.
Використовуючи цю функцію, додати можливість пошуку тексту в текстовому редакторі та вивести на екран усі
відповідні результати.
*/

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
	for i, row := range n.text {
		if reflect.DeepEqual(row, inputRow) {
			foundRow = true
			fmt.Println("Row number for your search query is --> ", i+1)
		}
	}

	if foundRow != true {
		fmt.Println("We could not found the row you have been searching for. Try one more time!")
	}
}

func main() {
	noteBookText := [][]string{}
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		text := scanner.Text()

		if text == "" {
			break
		}

		sliceText := strings.Split(text, " ")
		noteBookText = append(noteBookText, sliceText)
	}

	mainNotebook := Notebook{
		noteBookText,
	}

	fmt.Println("Text inside Notebook:")
	fmt.Println("******************************")
	fmt.Println(mainNotebook.GetTextNotebook())
	fmt.Println("******************************")

	fmt.Println("Here is the results of Search for Row inside Notebook Text:")
	testInput := []string{"Hello", "World"} // To test if SearchForNotebookRow works fine
	mainNotebook.SearchForNotebookRow(testInput)

	fmt.Println("Input a word combination you are looking for:")
	fmt.Scan(&inputStringForSearch)
	mainNotebook.SearchForNotebookText(inputStringForSearch)

}

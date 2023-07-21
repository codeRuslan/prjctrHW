package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type Action interface {
	PerformAction() string
}

type TextAction struct {
	text string
}

func (ta *TextAction) PerformAction() string {
	return ta.text
}

type DeleteWordAction struct {
	Action
	wordToDelete string
}

func (dwa *DeleteWordAction) PerformAction() string {
	text := dwa.Action.PerformAction()
	words := strings.Fields(text)
	var result []string

	for _, w := range words {
		if w != dwa.wordToDelete {
			result = append(result, w)
		}
	}

	return strings.Join(result, " ")
}

type InterchangeWordAction struct {
	Action
	wordToReplace   string
	replacementWord string
}

func (iwa *InterchangeWordAction) PerformAction() string {
	text := iwa.Action.PerformAction()
	words := strings.Fields(text)
	var result []string

	for _, w := range words {
		if w == iwa.wordToReplace {
			result = append(result, iwa.replacementWord)
		} else {
			result = append(result, w)
		}
	}

	return strings.Join(result, " ")
}

type CapitalizeDecorator struct {
	Action
}

func (cd *CapitalizeDecorator) PerformAction() string {
	text := cd.Action.PerformAction()
	return strings.ToUpper(text)
}

type DoubleSpaceDecorator struct {
	Action
}

// PerformAction подвоює всі пробіли у тексті.
func (dsd *DoubleSpaceDecorator) PerformAction() string {
	text := dsd.Action.PerformAction()
	return strings.ReplaceAll(text, " ", "  ")
}

func main() {
	pathFlag := flag.String("file", "text.txt", "Read file")
	useInterchangeFlag := flag.Bool("useInterchangeFlag", false, "Using interchange")
	useDeleteFlag := flag.Bool("useDeleteFlag", true, "Using delete")
	capitalizeDecoratorFlag := flag.Bool("useCapitalize", true, "Using capitalize")
	doubleSpaceDecoratorFlag := flag.Bool("useDoubleSpace", false, "Using Double Space")
	overwriteFileFlag := flag.Bool("overwriteFile", true, "Using Overwrite")
	flag.Parse()

	file, err := os.Open(*pathFlag)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	defer file.Close()

	content, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fileContent := string(content)

	var action Action = &TextAction{text: fileContent}

	fmt.Println("Initial input ->")
	fmt.Println(fileContent)

	fmt.Println("--------------------------------")
	fmt.Println("--------------------------------")

	if *useInterchangeFlag {
		newWord := "replacement"
		action = &InterchangeWordAction{Action: action, wordToReplace: "world", replacementWord: newWord}
	} else if *useDeleteFlag {
		action = &DeleteWordAction{Action: action, wordToDelete: "world"}
	} else {
		fmt.Println("No action selected. Please use -useInterchangeFlag or -useDeleteFlag flag.")
		return
	}

	if *capitalizeDecoratorFlag {
		action = &CapitalizeDecorator{Action: action}
	}

	if *doubleSpaceDecoratorFlag {
		action = &DoubleSpaceDecorator{Action: action}
	}

	outputText := action.PerformAction()

	fmt.Println("Output ->")
	fmt.Println(outputText)

	if *overwriteFileFlag {
		overwriteFile(pathFlag, outputText)
	}

}

func overwriteFile(filePath *string, text string) error {
	err := ioutil.WriteFile(*filePath, []byte(text), 0666)
	if err != nil {
		return fmt.Errorf("error writing to the file: %w", err)
	}
	fmt.Println("File overwritten successfully.")
	return nil
}

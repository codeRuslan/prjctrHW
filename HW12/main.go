package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type AdditionalDecorators struct {
	Actions
}

func (ad *AdditionalDecorators) doubleSpaceEverything() string {
	text := strings.ReplaceAll(ad.justPrint(), " ", "  ")
	return text
}

func (ad *AdditionalDecorators) capitalizeWord() string {
	text := ad.justPrint()
	text = strings.ToUpper(text)
	return text
}

type Actions interface {
	delete(word string)
	interchange(word string, newWord string)
	justPrint() string
}

type basicText struct {
	text string
}

func (bt *basicText) delete(word string) {
	words := strings.Fields(bt.text)
	var result []string

	for _, w := range words {
		if w != word {
			result = append(result, w)
		}
	}

	bt.text = strings.Join(result, " ")
}

func (bt *basicText) justPrint() string {
	return bt.text
}

func (bt *basicText) interchange(word string, newWord string) {
	words := strings.Fields(bt.text)
	var result []string

	for _, w := range words {
		if w == word {
			result = append(result, newWord)
		} else {
			result = append(result, w)
		}
	}

	bt.text = strings.Join(result, " ")
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

	var input = &basicText{
		text: fileContent,
	}

	fmt.Println("Initial input ->")
	fmt.Println(fileContent)

	fmt.Println("--------------------------------")
	fmt.Println("--------------------------------")

	if *useInterchangeFlag {
		newWord := "replacement"
		input.interchange("world", newWord)
		fmt.Println("Interchanged output ->")
		fmt.Println(input.text)

		if *capitalizeDecoratorFlag {
			decorator := AdditionalDecorators{Actions: input}
			outputText := decorator.capitalizeWord()
			fmt.Println("Capitalized output ->")
			fmt.Println(outputText)
			if *overwriteFileFlag {
				overwriteFile(pathFlag, outputText)
			}
		}
		if *doubleSpaceDecoratorFlag {
			decorator := AdditionalDecorators{Actions: input}
			outputText := decorator.doubleSpaceEverything()
			fmt.Println("Double Spaced output ->")
			fmt.Println(outputText)
			if *overwriteFileFlag {
				overwriteFile(pathFlag, outputText)
			}
		}
	} else if *useDeleteFlag {
		input.delete("world")
		fmt.Println("Deleted output ->")
		fmt.Println(input.text)

		if *capitalizeDecoratorFlag {
			decorator := AdditionalDecorators{Actions: input}
			outputText := decorator.capitalizeWord()
			fmt.Println("Capitalized output ->")
			fmt.Println(outputText)
			if *overwriteFileFlag {
				overwriteFile(pathFlag, outputText)
			}
		}
		if *doubleSpaceDecoratorFlag {
			decorator := AdditionalDecorators{Actions: input}
			outputText := decorator.doubleSpaceEverything()
			fmt.Println("Double Spaced output ->")
			fmt.Println(outputText)
			if *overwriteFileFlag {
				overwriteFile(pathFlag, outputText)
			}
		}

	} else {
		fmt.Println("No action selected. Please use -useInterchangeFlag or -useDeleteFlag flag.")
	}
}

func overwriteFile(path *string, text string) {
	err := ioutil.WriteFile(*path, []byte(text), 0666)
	if err != nil {
		fmt.Println("Error writing to the file:", err)
		return
	}
	fmt.Println("File overwritten successfully.")

}

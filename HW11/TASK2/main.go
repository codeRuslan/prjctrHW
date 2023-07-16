package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

func main() {
	fmt.Println("Please enter the template number:")
	fmt.Println("1. Words that begin with vowels")
	fmt.Println("2. Words that begin with consonants")
	fmt.Println("3. Words that begin with the same two letters")
	fmt.Println("4. Words that end with the same two letters")

	reader := bufio.NewReader(os.Stdin)
	templateNumStr, _ := reader.ReadString('\n')
	templateNum := strings.TrimSpace(templateNumStr)

	content, err := ioutil.ReadFile("text.txt")
	if err != nil {
		panic(err)
	}

	var regexStr string
	switch templateNum {
	case "1":
		regexStr = `\b[aeiouAEIOU][a-zA-Z]*\b`
	case "2":
		regexStr = `\b[^aeiouAEIOU][a-zA-Z]*\b`
	case "3":
		regexStr = `\b([a-zA-Z])\1[a-zA-Z]*\b`
	case "4":
		regexStr = `\b[a-zA-Z]*([a-zA-Z])\1\b`
	default:
		fmt.Println("Invalid template number")
		return
	}
	r := regexp.MustCompile(regexStr)
	matches := r.FindAllString(string(content), -1)
	for _, match := range matches {
		fmt.Println(match)
	}
}

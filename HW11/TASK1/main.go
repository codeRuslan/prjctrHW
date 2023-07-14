package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"regexp"
	"strings"
)

type Phone struct {
	ID          int
	PhoneNumber string
}

func main() {
	listPhones := []Phone{
		Phone{PhoneNumber: "1234567890"},
		Phone{PhoneNumber: "(123) 456-7890"},
		Phone{PhoneNumber: "(123)456-7890"},
		Phone{PhoneNumber: "123-456-7890"},
		Phone{PhoneNumber: "123.456.7890"},
		Phone{PhoneNumber: "123 456 7890"},
	}
	db, err := gorm.Open("postgres", "user=ruslanpilipyuk password=rus090203 dbname=postgres sslmode=disable")
	if err != nil {
		panic(err.Error())
	}
	// Close the connection when the database transaction is over.
	defer db.Close() // Not required!

	db.DropTable(&Phone{})
	db.CreateTable(&Phone{})

	for _, i := range listPhones {
		db.Create(&i)
	}

	var initPhoneNumbers []Phone
	db.Find(&initPhoneNumbers)
	fmt.Println("Phone Number Database:")
	for _, phone := range initPhoneNumbers {
		fmt.Println(phone.PhoneNumber)
	}

	regex := `\+?\d{0,2}[-.\s(]?\d{1,3}[-.\s)]?\d{1,3}[-.\s]?\d{1,4}`
	searchFormat := "(###) ###-####"
	searchPhoneNumbers := searchPhoneNumbersInFormat(listPhones, regex, searchFormat)
	fmt.Printf("\nPhone Numbers in Format %s:\n", searchFormat)
	for _, phoneNumber := range searchPhoneNumbers {
		fmt.Println(phoneNumber)
	}
}
func applyRegex(phoneNumber string) string {
	regex := regexp.MustCompile(`\D`)
	return regex.ReplaceAllString(phoneNumber, "")
}

func searchPhoneNumbersInFormat(phones []Phone, regex, format string) []string {
	var matchedPhoneNumbers []string
	searchRegex := formatToRegex(regex, format)

	for _, phone := range phones {
		if searchRegex.MatchString(phone.PhoneNumber) {
			matchedPhoneNumbers = append(matchedPhoneNumbers, phone.PhoneNumber)
		}
	}

	return matchedPhoneNumbers
}

func formatToRegex(regex, format string) *regexp.Regexp {
	format = regexp.QuoteMeta(format)
	format = regexp.MustCompile(`\\#`).ReplaceAllLiteralString(format, `\d`)

	regex = fmt.Sprintf("^%s$", regex)
	regex = regexp.MustCompile(`\(\?i\)`).ReplaceAllLiteralString(regex, "")

	return regexp.MustCompile(strings.ReplaceAll(regex, "###", format))
}

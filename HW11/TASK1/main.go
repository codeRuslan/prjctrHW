package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"regexp"
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
	defer db.Close()

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

	// Регулярний вираз для пошуку телефонних номерів у різних форматах
	phoneRegex := regexp.MustCompile(`\(?(\d{3})\)?[-.\s]?(\d{3})[-.\s]?(\d{4})`)

	fmt.Println("\nMatching Phone Numbers:")
	for _, phone := range initPhoneNumbers {
		if phoneRegex.MatchString(phone.PhoneNumber) {
			fmt.Println(phone.PhoneNumber)
		}
	}
}

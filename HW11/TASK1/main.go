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
	db, err := gorm.Open("postgres", "user=ruslanpilipyuk password=XXXXX dbname=postgres sslmode=disable")
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
	fmt.Println("Initial Phone Numbers")
	for _, phone := range initPhoneNumbers {
		fmt.Println(phone.PhoneNumber)
	}

	var phones []Phone
	db.Find(&phones)

	for i := range phones {
		phones[i].PhoneNumber = applyRegex(phones[i].PhoneNumber)
		db.Save(&phones[i])
	}

	var updatedPhones []Phone
	db.Find(&updatedPhones)
	fmt.Println("Updated Phone Numbers")
	for _, phone := range updatedPhones {
		fmt.Println(phone.PhoneNumber)
	}
}

func applyRegex(phoneNumber string) string {
	regex := regexp.MustCompile(`\D`)
	return regex.ReplaceAllString(phoneNumber, "")
}

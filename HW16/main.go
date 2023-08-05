package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
)

type userData struct {
	ID   int
	Name string
}

type userInfoData struct {
	ID             int
	PassportNumber string
}

type discountInfoData struct {
	ID                     int
	AmountOfDiscountPoints int
}

func main() {
	config := mysql.Config{
		User:   "user",
		Passwd: "password",
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "my-app",
	}

	db, err := sql.Open("mysql", config.FormatDSN())
	if err != nil {
		log.Fatal(err.Error())
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err.Error())
	}

	usersData := []userData{
		{ID: 1, Name: "Elizabhet"},
		{ID: 2, Name: "John"},
	}

	userInfoData := []userInfoData{
		{ID: 1, PassportNumber: "JH21313"},
		{ID: 2, PassportNumber: "KK23132"},
	}

	discountInfoData := []discountInfoData{
		{ID: 1, AmountOfDiscountPoints: 201},
		{ID: 2, AmountOfDiscountPoints: 13},
	}

	insertOrUpdateData(db, "userstable", usersData)
	insertOrUpdateData(db, "user_info", userInfoData)
	insertOrUpdateData(db, "discount_info", discountInfoData)

	retrieveValue(db, "userstable")
	retrieveValue(db, "user_info")
	retrieveValue(db, "discount_info")
}

func insertOrUpdateData(db *sql.DB, tableName string, data interface{}) {
	switch tableName {
	case "userstable":
		userstableData, ok := data.([]userData)
		if !ok {
			log.Fatal("Invalid data type for userstableData")
		}
		for _, data := range userstableData {
			_, err := db.Exec("INSERT INTO userstable (id, name) VALUES (?, ?) ON DUPLICATE KEY UPDATE name = VALUES(name)", data.ID, data.Name)
			if err != nil {
				log.Fatal(fmt.Errorf("failed to insert/update row for userstable: %w", err))
			}
		}
	case "user_info":
		userInfoData, ok := data.([]userInfoData)
		if !ok {
			log.Fatal("Invalid data type for userInfoData")
		}
		for _, data := range userInfoData {
			_, err := db.Exec("INSERT INTO user_info (id, passport_number) VALUES (?, ?) ON DUPLICATE KEY UPDATE passport_number = VALUES(passport_number)", data.ID, data.PassportNumber)
			if err != nil {
				log.Fatal(fmt.Errorf("failed to insert/update row for user_info: %w", err))
			}
		}
	case "discount_info":
		discountInfoData, ok := data.([]discountInfoData)
		if !ok {
			log.Fatal("Invalid data type for discountInfoData")
		}
		for _, data := range discountInfoData {
			_, err := db.Exec("INSERT INTO discount_info (id, amount_of_discount_points) VALUES (?, ?) ON DUPLICATE KEY UPDATE amount_of_discount_points = VALUES(amount_of_discount_points)", data.ID, data.AmountOfDiscountPoints)
			if err != nil {
				log.Fatal(fmt.Errorf("failed to insert/update row for discount_info: %w", err))
			}
		}
	default:
		log.Fatal(fmt.Errorf("unsupported table name: %s", tableName))
	}
}

func retrieveValue(db *sql.DB, tableName string) {
	rows, err := db.Query(fmt.Sprintf("SELECT * FROM %s", tableName))
	if err != nil {
		log.Fatal(fmt.Errorf("failed to retrieve values from %s: %w", tableName, err))
	}
	defer rows.Close()

	fmt.Printf("Values from table %s:\n", tableName)
	var (
		id        int
		str_input string
	)
	for rows.Next() {
		err := rows.Scan(&id, &str_input)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(id, str_input)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(fmt.Errorf("error while iterating rows for %s: %w", tableName, err))
	}
}

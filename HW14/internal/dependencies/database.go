package dependencies

import (
	"awesomeProject3/internal/types"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

func InitialMigration() {
	db, err := gorm.Open("sqlite3", Cfg.Database.DatabaseLocation)
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect database")
	}
	defer db.Close()

	db.LogMode(true)
	if err := db.AutoMigrate(&types.User{}).Error; err != nil {
		fmt.Println("Error during auto-migration:", err.Error())
		panic("Failed to auto-migrate database")
	}

	fmt.Println("Database has been successfully migrated")
}

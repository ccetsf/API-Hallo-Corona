package database

import (
	"fmt"
	"hallo-corona/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DatabaseInit() {
	var err error

	// data source name
	dsn := "root:@tcp(localhost:3306)/halco?charset=utf8mb4&parseTime=True&loc=Local"
	// connection to database
	DB, err = gorm.Open(mysql.Open(dsn))
	if err != nil {
		fmt.Println(err)
		panic("Failed to connect to database")
	}

	//migration models to database
	err = DB.AutoMigrate(
		&models.User{},
	)
	if err != nil {
		fmt.Println(err)
		panic("Failed migration models to database")
	}

	fmt.Println("Success to connect and migrate to database")
}

package utils

import (
	"expense/models"
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
)

//ConnectDB function
func ConnectDB() *gorm.DB {

	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")

	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password) //Build connection string
	fmt.Println(dbURI)

	db, err := gorm.Open("postgres", dbURI)

	// Migrate the schema
	db.AutoMigrate(
		&models.Expense{},
		&models.User{})

	//close db when not in use
	defer db.Close()
	if err != nil {
		fmt.Println("error=%s", err)
	}

	fmt.Println("connection=%s", db)
	return db
}

package main

// package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"expense/routes"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

func main() {
	e := godotenv.Load()

	if e != nil {
		fmt.Print(e)
	}
	fmt.Println(e)
	godotenv.Load()

	port := os.Getenv("PORT")
	fmt.Println(port)

	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")

	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password) //Build connection string
	fmt.Println(dbURI)

	conn, err := gorm.Open("postgres", dbURI)
	if err != nil {
		fmt.Print(err)
	}

	fmt.Println("connection=%s", conn)

	// Handle routes
	http.Handle("/", routes.Handlers())
	// http.ListenAndServe()

	// serve
	log.Printf("Server up on port '%s'", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))

}

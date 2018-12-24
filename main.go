package main

// package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/linux08/expense/routes"
)

func main() {
	e := godotenv.Load()

	if e != nil {
		fmt.Print(e)
	}
	godotenv.Load()

	port := os.Getenv("PORT")
	fmt.Println(port)

	// Handle routes
	http.Handle("/", routes.Handlers())
	// http.ListenAndServe()

	// serve
	log.Printf("Server up on port '%s'", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))

}

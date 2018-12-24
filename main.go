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

// // Handlers : Use gorilla mux.
// // CRUD + Vote Handlers.
// func Handlers() *mux.Router {
// 	r := mux.NewRouter().StrictSlash(true)

// 	r.HandleFunc("/api", testApi).Methods("GET")
// 	r.HandleFunc("/api/expenses", getExpenses).Methods("GET")
// 	r.HandleFunc("/api/expenses/{id}", getExpense).Methods("GET")
// 	r.HandleFunc("/api/expenses", createExpense).Methods("POST")
// 	r.HandleFunc("/api/expenses/{id}", updateExpense).Methods("PUT")
// 	r.HandleFunc("/api/expenses/{id}", deleteExpense).Methods("DELETE")
// 	return r
// }

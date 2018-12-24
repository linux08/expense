package routes

import "github.com/gorilla/mux"

// Handlers : Use gorilla mux.
// CRUD + Vote Handlers.
func Handlers() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/api", testApi).Methods("GET")
	r.HandleFunc("/api/expenses", getExpenses).Methods("GET")
	r.HandleFunc("/api/expenses/{id}", getExpense).Methods("GET")
	r.HandleFunc("/api/expenses", createExpense).Methods("POST")
	r.HandleFunc("/api/expenses/{id}", updateExpense).Methods("PUT")
	r.HandleFunc("/api/expenses/{id}", deleteExpense).Methods("DELETE")
	return r
}

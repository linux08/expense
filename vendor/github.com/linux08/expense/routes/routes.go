package routes

import (
	"github.com/gorilla/mux"
	"github.com/linux08/expense/controllers"
)

// Handlers : Use gorilla mux.
// CRUD + Vote Handlers.
func Handlers() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/api", controllers.testAPI).Methods("GET")
	r.HandleFunc("/api/expenses", controllers.getExpenses).Methods("GET")
	r.HandleFunc("/api/expenses/{id}", controllers.getExpense).Methods("GET")
	r.HandleFunc("/api/expenses", controllers.createExpense).Methods("POST")
	r.HandleFunc("/api/expenses/{id}", controllers.updateExpense).Methods("PUT")
	r.HandleFunc("/api/expenses/{id}", controllers.deleteExpense).Methods("DELETE")
	return r
}

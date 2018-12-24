package routes

import (
	"expense/models"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/linux08/expense/controllers"
)

//Init expense as a slice of expense struct
var expense []models.Expense

// Handlers : Use gorilla mux.
// CRUD + Vote Handlers.
func Handlers() *mux.Router {

	r := mux.NewRouter().StrictSlash(true)
	r.Use(commonMiddleware)

	r.HandleFunc("/api", controllers.TestAPI).Methods("GET")
	r.HandleFunc("/api/expenses", controllers.GetExpenses).Methods("GET")
	r.HandleFunc("/api/expenses/{id}", controllers.GetExpense).Methods("GET")
	r.HandleFunc("/api/expenses", controllers.CreateExpense).Methods("POST")
	r.HandleFunc("/api/expenses", controllers.UpdateExpenses).Methods("PUT")
	r.HandleFunc("/api/expenses/{id}", controllers.UpdateExpense).Methods("PUT")
	r.HandleFunc("/api/expenses/{id}", controllers.DeleteExpense).Methods("DELETE")
	return r
}

//set content-type
func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

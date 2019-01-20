package routes

import (
	"expense/models"
	"net/http"

	"expense/controllers"

	"github.com/gorilla/mux"
)

//Init expense as a slice of expense struct
var expense []models.Expense

// Handlers : Use gorilla mux.
func Handlers() *mux.Router {

	r := mux.NewRouter().StrictSlash(true)
	// utils.parseClaims()
	r.Use(commonMiddleware)
	// r.Use(auth.JwtVerify)

	r.HandleFunc("/api", controllers.TestAPI).Methods("GET")
	r.HandleFunc("/api/expenses", controllers.GetExpenses).Methods("GET")
	r.HandleFunc("/api/expenses/{id}", controllers.GetExpense).Methods("GET")
	r.HandleFunc("/api/expenses", controllers.CreateExpense).Methods("POST")
	r.HandleFunc("/api/expenses/{id}", controllers.UpdateExpense).Methods("PUT")
	r.HandleFunc("/api/expenses/{id}", controllers.DeleteExpense).Methods("DELETE")

	r.HandleFunc("/api/register", controllers.CreateUser).Methods("POST")
	r.HandleFunc("/api/login", controllers.Login).Methods("POST")
	r.HandleFunc("/api/user", controllers.FetchUsers).Methods("GET")

	r.HandleFunc("/api/user/{id}", controllers.GetUser).Methods("GET")
	r.HandleFunc("/api/user/{id}", controllers.UpdateUser).Methods("PUT")
	r.HandleFunc("/api/user/{id}", controllers.DeleteUser).Methods("DELETE")

	return r
}

//set content-type
func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

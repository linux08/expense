package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/linux08/expense/models"
)

//Init expense as a slice of expense struct
var expense []models.Expense

type exp struct{}
type expenseArray []models.Expense

//intialized sample objec
func InitSampleObj() (expenses expenseArray) {
	//Init expensest  var as a slice of expense struct

	expense := make([]models.Expense, 0, 5)

	expense = append(expense,
		models.Expense{ID: "1", Name: "good", Reason: "i need to save", Vat: "132",
			User: &models.User{ID: "13", Name: "linx08", Gender: "male", Email: "abimbola120@gmail.com", Password: "pass123"}})

	expense = append(expense,
		models.Expense{ID: "2", Name: "godod", Reason: "i need to save", Vat: "13f2",
			User: &models.User{ID: "13", Name: "linx08", Gender: "male", Email: "abimbola120@gmail.com", Password: "pass123"}})
	return expense
}

func GetExpenses(w http.ResponseWriter, r *http.Request) {
	var expense = InitSampleObj()
	json.NewEncoder(w).Encode(expense)
}

func GetExpense(w http.ResponseWriter, r *http.Request) {
	// fmt.Println("dddjj")
	w.Write([]byte("Hello World"))
}

func CreateExpense(w http.ResponseWriter, r *http.Request) {
	// fmt.Println("dddjj")
	w.Write([]byte("Hello World"))
}

func UpdateExpense(w http.ResponseWriter, r *http.Request) {
	// fmt.Println("dddjj")
	w.Write([]byte("Hello World"))
}

func DeleteExpense(w http.ResponseWriter, r *http.Request) {
	// fmt.Println("dddjj")
	w.Write([]byte("Hello World"))
}

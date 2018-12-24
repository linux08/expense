package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/linux08/expense/models"
)

//Init expense as a slice of expense struct
var expense []models.Expense

type expenses []models.Expense

//intialized sample objec
func InitSampleObj() expenses {
	//Init expensest  var as a slice of expense struct
	expense := append(expense,
		models.Expense{ID: "1", Name: "good", Reason: "i need to save", Vat: "132",
			User: &models.User{ID: "13", Name: "linx08", Gender: "male", Email: "abimbola120@gmail.com", Password: "pass123"}})
	return expense
}

func GetExpenses(w http.ResponseWriter, r *http.Request) {
	// fmt.Println("dddjj")
	var expense = InitSampleObj()
	fmt.Println("ed", expense)
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

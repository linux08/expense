package controllers

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/linux08/expense/models"
)

//Init expense as a slice of expense struct
var expense []models.Expense

type exp struct{}
type expenseArray []models.Expense

//intialized sample object
func InitSampleObj() (expenses expenseArray) {
	//Init expenses  var as a slice of expense struct

	expense := make([]models.Expense, 0, 5)

	expense = append(expense,
		models.Expense{ID: "1", Name: "good", Reason: "i need to save", Vat: "132",
			User: &models.User{ID: "13", Name: "linx08", Gender: "male", Email: "abimbola120@gmail.com", Password: "pass123"}})

	expense = append(expense,
		models.Expense{ID: "2", Name: "godod", Reason: "i need to save", Vat: "13f2",
			User: &models.User{ID: "15", Name: "linx08", Gender: "male", Email: "abimbola120@gmail.com", Password: "pass123"}})
	return expense
}

func GetExpenses(w http.ResponseWriter, r *http.Request) {
	var expense = InitSampleObj()
	json.NewEncoder(w).Encode(expense)
}

func GetExpense(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	expense := InitSampleObj()
	for _, item := range expense {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&models.Expense{})
}

func CreateExpense(w http.ResponseWriter, r *http.Request) {
	var expense models.Expense
	_ = json.NewDecoder(r.Body).Decode(&expense)
	expense.ID = strconv.Itoa(rand.Intn(1000))
	var expenses = InitSampleObj()
	expenses = append(expenses, expense)
	fmt.Println(expense)
	json.NewEncoder(w).Encode(expense)
}

func UpdateExpenses(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	expense := InitSampleObj()
	for index, item := range expense {
		if item.ID == params["id"] {
			expense = append(expense[:index], expense[index+1:]...)
			var expenseObj models.Expense
			_ = json.NewDecoder(r.Body).Decode(&expenseObj)
			expenseObj.ID = strconv.Itoa(rand.Intn(1000))
			var expenses = InitSampleObj()
			expenses = append(expenses, expenseObj)
			fmt.Println(expenses)
			json.NewEncoder(w).Encode(expenses)
		}
	}
}

func UpdateExpense(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	expense := InitSampleObj()
	for index, item := range expense {
		if item.ID == params["id"] {
			expense = append(expense[:index], expense[index+1:]...)
			var expenseObj models.Expense
			_ = json.NewDecoder(r.Body).Decode(&expenseObj)
			expenseObj.ID = strconv.Itoa(rand.Intn(1000))
			expense = append(expense, expenseObj)
			fmt.Println(expense)
			json.NewEncoder(w).Encode(expense)
		}
	}
}

func DeleteExpense(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	expense := InitSampleObj()
	for index, item := range expense {
		if item.ID == params["id"] {
			expense = append(expense[:index], expense[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(expense)
}

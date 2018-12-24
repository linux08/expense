package controllers

import (
	"encoding/json"
	"net/http"
		"github.com/linux08/expense/models"
)
//Init expense as a slice of expense struct
var expense []models.Expense

func InitSampleObj(){
	//Init expenses  var as a slice of expense struct
expense =append(expense,
	Expense{ ID: "1",Name: "good", Reason: "i need to save",Vat: "132",
	User: &User{ ID: "13",Name: "linx08", Gender: "male",Email: "abimbola120@gmail.com",Password: "pass123" }
})
return expense
}


func GetExpenses(w http.ResponseWriter, r *http.Request) {
	// fmt.Println("dddjj")
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

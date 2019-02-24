package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"expense/models"

	"github.com/gorilla/mux"
)

//Init expense as a slice of expense struct
var expense []models.Expense

type exp struct{}

type expenseArray []models.Expense

// var db = utils.ConnectDB()

//intialized sample object
func InitSampleObj() (expenses expenseArray) {
	//Init expenses  var as a slice of expense struct

	expense := make([]models.Expense, 0, 5)

	expense = append(expense,
		models.Expense{Name: "good", Reason: "i need to save", Vat: "132"}) //  User: &models.User{Name: "linx08", Gender: "male", Email: "abimbola120@gmail.com", Password: "pass123"}

	expense = append(expense,
		models.Expense{Name: "godod", Reason: "i need to save", Vat: "13f2"}) //  User: &models.User{Name: "linx08", Gender: "male", Email: "abimbola120@gmail.com", Password: "pass123"}

	return expense
}

func GetExpenses(w http.ResponseWriter, r *http.Request) {
	var expenses []models.Expense
	db.Find(&expenses)

	json.NewEncoder(w).Encode(expenses)
}

//Gte expense that belongs to a user
func GetExpenseForAUser(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user")

	fmt.Println("the user = struct", user)
	userInfo := models.Token{}
	bodyBytes, _ := json.Marshal(user)
	json.Unmarshal(bodyBytes, &userInfo)

	// expense := &models.Expense{}
	var expense []models.Expense
	// exp :=
	db.Where("user_id = ?", userInfo.UserID).Find(&expense)
	json.NewEncoder(w).Encode(expense)
}

func GetExpense(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	expense := &models.Expense{}
	exp := db.Where("ID = ?", params["id"]).First(expense).Preload("User")
	json.NewEncoder(w).Encode(exp)
}

func CreateExpense(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user")

	fmt.Println("the user = struct", user)
	userInfo := models.Token{}
	bodyBytes, _ := json.Marshal(user)
	json.Unmarshal(bodyBytes, &userInfo)
	// fmt.Println(userInfo.UserID)

	expense := &models.Expense{}

	json.NewDecoder(r.Body).Decode(expense)

	// fmt.Println(user["Name"])
	expense.UserID = userInfo.UserID
	// expense.user_id = userInfo.UserID
	fmt.Print("expens", expense)
	createdExpense := db.Create(expense)
	// cExpense := db.Preload("Expense").Find(&models.User{})
	var errMessage = createdExpense.Error
	if createdExpense.Error != nil {
		fmt.Println(errMessage)
	}
	json.NewEncoder(w).Encode(createdExpense)

	// var expense models.Expense
	// _ = json.NewDecoder(r.Body).Decode(&expense)
	// // expense.ID = strconv.Itoa(rand.Intn(1000))
	// var expenses = InitSampleObj()
	// expenses = append(expenses, expense)
	// fmt.Println(expense)
	// json.NewEncoder(w).Encode(expense)
}

func UpdateExpenses(w http.ResponseWriter, r *http.Request) {
	// params := mux.Vars(r)
	expense := InitSampleObj()
	// for index, item := range expense {
	// 	if item.ID == params["id"] {
	// 		expense = append(expense[:index], expense[index+1:]...)
	// 		var expenseObj models.Expense
	// 		_ = json.NewDecoder(r.Body).Decode(&expenseObj)
	// 		// expenseObj.ID = strconv.Itoa(rand.Intn(1000))
	// 		var expenses = InitSampleObj()
	// 		expenses = append(expenses, expenseObj)
	// 		fmt.Println(expenses)
	// 		json.NewEncoder(w).Encode(expenses)
	// 	}
	// }
	json.NewEncoder(w).Encode(expense)
}

func UpdateExpense(w http.ResponseWriter, r *http.Request) {
	// params := mux.Vars(r)
	expense := InitSampleObj()
	// for index, item := range expense {
	// if item.ID == params["id"] {
	// 	expense = append(expense[:index], expense[index+1:]...)
	// 	var expenseObj models.Expense
	// 	_ = json.NewDecoder(r.Body).Decode(&expenseObj)
	// 	expenseObj.ID = strconv.Itoa(rand.Intn(1000))
	// 	expense = append(expense, expenseObj)
	// 	fmt.Println(expense)
	// 	json.NewEncoder(w).Encode(expense)
	// }

	json.NewEncoder(w).Encode(expense)
}

func DeleteExpense(w http.ResponseWriter, r *http.Request) {
	// params := mux.Vars(r)
	expense := InitSampleObj()
	// for index, item := range expense {
	// 	if item.ID == params["id"] {
	// 		expense = append(expense[:index], expense[index+1:]...)
	// 		break
	// 	}
	// }
	json.NewEncoder(w).Encode(expense)
}

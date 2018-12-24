package controllers

import (
	"net/http"
)

func GetExpenses(w http.ResponseWriter, r *http.Request) {
	// fmt.Println("dddjj")
	w.Write([]byte("Hello World"))
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

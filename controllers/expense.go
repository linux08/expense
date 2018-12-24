package controllers

import (
	"net/http"
)

func getExpenses(w http.ResponseWriter, r *http.Request) {
	// fmt.Println("dddjj")
	w.Write([]byte("Hello World"))
}

func getExpense(w http.ResponseWriter, r *http.Request) {
	// fmt.Println("dddjj")
	w.Write([]byte("Hello World"))
}

func createExpense(w http.ResponseWriter, r *http.Request) {
	// fmt.Println("dddjj")
	w.Write([]byte("Hello World"))
}

func updateExpense(w http.ResponseWriter, r *http.Request) {
	// fmt.Println("dddjj")
	w.Write([]byte("Hello World"))
}

func deleteExpense(w http.ResponseWriter, r *http.Request) {
	// fmt.Println("dddjj")
	w.Write([]byte("Hello World"))
}

package controllers

import (
	"encoding/json"
	"expense/models"
	"expense/utils"
	"fmt"
	"net/http"
)

var db = utils.ConnectDB()

func TestAPI(w http.ResponseWriter, r *http.Request) {
	// fmt.Println("API live and kicking")
	w.Write([]byte("API live and kicking"))
}

//CreateUser function
func CreateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("jjdjdjjd")
	user := &models.User{}
	// fmt.Println(json.NewDecoder(r.Body))
	json.NewDecoder(r.Body).Decode(user)
	db.NewRecord(user)

	var createdUser = db.Create(user)
	json.NewEncoder(w).Encode(createdUser)
}

//FetchUser function
func FetchUsers(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	db.Find(&users)
	json.NewEncoder(w).Encode(users)
}

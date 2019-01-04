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
	var user models.User
	// fmt.Println(json.NewDecoder(r.Body))
	_ = json.NewDecoder(r.Body).Decode(&user)
	// reqBody := json.NewDecoder(r.Body).Decode(&models.User)
	fmt.Println(user)
	db.Create(user)
	// fmt.Println(users)
	// db.Create()
	json.NewEncoder(w).Encode(user)
}

//FetchUser function
func FetchUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("fetch aall user")
	var user models.User
	// fmt.Println(json.NewDecoder(r.Body))
	// db.Find(&user)
	users := db.Find(&user)

	// db.Create()
	json.NewEncoder(w).Encode(users)
}
 
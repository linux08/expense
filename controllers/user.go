package controllers

import (
	"encoding/json"
	"expense/models"
	"expense/utils"
	"fmt"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

var db = utils.ConnectDB()

func TestAPI(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("API live and kicking"))
}

//CreateUser function
func CreateUser(w http.ResponseWriter, r *http.Request) {

	type ErrorResponse struct {
		Err string
	}
	user := &models.User{}
	pass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("err")
		err := ErrorResponse{
			Err: "Password validation failed",
		}
		json.NewEncoder(w).Encode(err)
	}
	json.NewDecoder(r.Body).Decode(user)
	user.Password = string(pass)

	// json.NewEncoder(w).Encode(user)
	// db.NewRecord(user)
	createdUser := db.Create(user)
	json.NewEncoder(w).Encode(createdUser)
}

//FetchUser function
func FetchUsers(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	db.Find(&users)
	json.NewEncoder(w).Encode(users)
}

package controllers

import (
	"encoding/json"
	"expense/models"
	"expense/utils"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

type ErrorResponse struct {
	Err string
}
type error interface {
	Error() string
}

var db = utils.ConnectDB()

func MagaAPI(w http.ResponseWriter, r *http.Request) {
	fmt.Println("i am here")
}

func TestAPI(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("API live and kicking"))
}

//CreateUser function
func CreateUser(w http.ResponseWriter, r *http.Request) {

	user := &models.User{}
	pass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
		err := ErrorResponse{
			Err: "Password validation failed",
		}
		json.NewEncoder(w).Encode(err)
	}
	json.NewDecoder(r.Body).Decode(user)
	user.Password = string(pass)

	createdUser := db.Create(user)
	fmt.Print("created-user")
	fmt.Println(createdUser.Error)
	var errMessage = createdUser.Error

	if createdUser.Error != nil {
		fmt.Println(errMessage)
	}
	json.NewEncoder(w).Encode(createdUser)
}

//FetchUser function
func FetchUsers(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	db.Find(&users)
	json.NewEncoder(w).Encode(users)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	params := mux.Vars(r)
	var id = params["id"]
	db.First(&user, id)
	json.NewDecoder(r.Body).Decode(user)
	db.Save(&user)
	json.NewEncoder(w).Encode(&user)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var id = params["id"]
	var user models.User
	db.First(&user, id)
	db.Delete(&user)
	json.NewEncoder(w).Encode("User deleted")
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var id = params["id"]
	var user models.User
	db.First(&user, id)
	json.NewEncoder(w).Encode(&user)
}

package controllers

import (
	"encoding/json"
	"expense/models"
	"expense/utils"
	"fmt"
	"net/http"
	"os"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
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

func Login(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	err := json.NewDecoder(r.Body).Decode(user) //decode the request body into struct and failed if any error occur
	if err != nil {
		var resp = map[string]interface{}{"status": false, "message": "Invalid request"}
		json.NewEncoder(w).Encode(resp)
		return
	}

	resp := FindOne(models.User{})
	json.NewEncoder(w).Encode(resp)
	// if err != nil {
	// 	var resp = map[string]interface{}{"status": false, "message": "Invalid request"}
	// 	json.NewEncoder(w).Encode(resp)
	// 	return
	// }
	// db.First(&user, id)
	// fmt.Println(userResp)
	// json.NewEncoder(w).Encode(userResp)

}

func FindOne(userObj models.User) map[string]interface{} {
	fmt.Print("got ehheh")
	user := &models.User{}
	email := userObj.Email
	err := db.Where(&models.User{Email: email}).First(user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			var resp = map[string]interface{}{"status": false, "message": "Email address not found"}
			// json.NewEncoder(w).Encode(resp)
			return resp
		}
		// var resp = map[string]interface{}{"status": false, "Connection error. Please retry"}
		// return json.NewEncoder(w).Encode(resp)
	}

	err = bcrypt.CompareHashAndPassword([]byte(userObj.Password), []byte(userObj.Password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword { //Password does not match!
		var resp = map[string]interface{}{"status": false, "message": "Invalid login credentials. Please try again"}
		// json.NewEncoder(w).Encode(resp)
		return resp
	}

	user.Password = ""

	//Create JWT token
	tk := &models.Token{UserID: user.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	var respObj map[string]interface{}
	// respObj := user
	respObj["Token"] = tokenString //Store the token in the response
	var resp = map[string]interface{}{"status": false, "message": "logged in"}
	resp["user"] = user
	return resp
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

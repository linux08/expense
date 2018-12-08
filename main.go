package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("ode")
	r := mux.NewRouter()
	r.HandleFunc("/", Login)

	http.Handle("/", r)
}

func Login(w http.ResponseWriter, r *http.Request) {

}

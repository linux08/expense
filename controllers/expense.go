package controllers

import (
	"fmt"
	"net/http"
)

func ListHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("dddjj")
	w.Write([]byte("Hello World"))
}

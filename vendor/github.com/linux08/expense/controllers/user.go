package controllers

import (
	"net/http"
)

func testAPI(w http.ResponseWriter, r *http.Request) {
	// fmt.Println("API live and kicking")
	w.Write([]byte("API live and kicking"))
}

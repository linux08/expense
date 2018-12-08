package utils

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/context"

	"gopkg.in/mgo.v2"
)

var (
	s          *mgo.Session
	err        error
	database   = "expense"
	collection = "expense"
)

//  Log all request
//  Apache access_log format
func LogHandler(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		nw := NewLogResponseWriter(w)
		f(nw, r)

		log.Println(nw.String(r))
	}
}

// Recover from error => log
func RecoverHandler(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				log.Println(fmt.Sprintf("Recover from '%v'", r))
			}
		}()
		f(w, r)
	}
}

// Get a new mongo session by request : "crash friendly".
// Close the session after using it.
func DBHandler(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		db := s.Copy()
		defer db.Close()
		context.Set(r, "db", db.DB("expense"))
		f(w, r)
	}
}

// Print the right headers when the method "OPTIONS" is called.
// Allow all origin (*)
func CorsHandler(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		if r.Method == "OPTIONS" {
			return
		}
		f(w, r)
	}
}

// Map all sub middlewares
func MiddlewareHandler(f http.HandlerFunc) http.HandlerFunc {
	return RecoverHandler(LogHandler(CorsHandler(DBHandler(f))))
}

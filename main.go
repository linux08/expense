package main

// package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"expense/controllers"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	e := godotenv.Load()

	if e != nil {
		fmt.Print(e)
	}
	godotenv.Load()

	port := os.Getenv("PORT")
	fmt.Println(port)

	// Handle routes
	http.Handle("/", Handlers())

	// serve
	log.Printf("Server up on port '%s'", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))

}

// Handlers : Use gorilla mux.
// CRUD + Vote Handlers.
func Handlers() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/api", controllers.ListHandler).Methods("GET")

	// r.HandleFunc("/api", utils.MiddlewareHandler(controllers.ListHandler)).Methods("GET")
	// r.HandleFunc("/api/v1/polls/{poll}", MiddlewareHandler(getHandler)).Methods("GET")
	// r.HandleFunc("/api/v1/polls", MiddlewareHandler(postHandler)).Methods("POST")
	// r.HandleFunc("/api/v1/polls/{poll}", MiddlewareHandler(putHandler)).Methods("PUT")
	// r.HandleFunc("/api/v1/polls/{poll}", MiddlewareHandler(deleteHandler)).Methods("DELETE")
	// r.HandleFunc("/api/v1/polls/{poll}/answers/{answer}", MiddlewareHandler(voteHandler)).Methods("POST")
	// r.NotFoundHandler = LogHandler(notFoundHandler)

	return r
}

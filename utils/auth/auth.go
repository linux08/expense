package auth

import (
	"context"
	"encoding/json"
	"expense/models"
	"fmt"
	"net/http"
	"os"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
)

// Middleware function, which will be called for each request
func JwtVerify(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		notAuth := []string{"/api", "/api/register", "/api/login"} //List of endpoints that doesn't require auth
		requestPath := r.URL.Path                                  //current request path

		//check if request does not need authentication, serve the request if it doesn't need it
		for _, value := range notAuth {

			if value == requestPath {
				next.ServeHTTP(w, r)
				return
			}
		}

		tokenHeader := r.Header.Get("Authorization") //Grab the token from the header

		if tokenHeader == "" {
			//Token is missing, returns with error code 403 Unauthorized
			var resp = map[string]interface{}{"status": false, "message": "Missing auth token"}
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode(resp)
			return
		}

		splitted := strings.Split(tokenHeader, " ") //The token normally comes in format `Bearer {token-body}`, we check if the retrieved token matched this requirement
		if len(splitted) != 2 {
			var resp = map[string]interface{}{"status": false, "message": "Invalid/Malformed auth token"}
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode(resp)
			return
		}

		tokenPart := splitted[1] //Grab the token part, what we are truly interested in
		tk := &models.Token{}

		token, err := jwt.ParseWithClaims(tokenPart, tk, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("token_password")), nil
		})

		if err != nil { //Malformed token, returns with http code 403 as usual
			var resp = map[string]interface{}{"status": false, "message": "Malformed authentication token"}
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode(resp)
			return
		}

		if !token.Valid { //Token is invalid, maybe not signed on this server
			var resp = map[string]interface{}{"status": false, "message": "Token is not valid"}
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode(resp)
			return
		}

		//Everything went well, proceed with the request and set the caller to the user retrieved from the parsed token
		fmt.Sprintf("User %", tk.UserID) //Useful for monitoring
		ctx := context.WithValue(r.Context(), "user", tk.UserID)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r) //proceed in the middleware chain!
	})
}

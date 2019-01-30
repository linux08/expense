package auth

import (
	"context"
	"encoding/json"
	"expense/models"
	"fmt"
	"net/http"
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

		var authHdr = r.Header.Get("Authorization") //Grab the token from the header

		authHdr = strings.TrimSpace(authHdr)
		splitted := strings.Split(authHdr, ".")

		if authHdr == "" {
			//Token is missing, returns with error code 403 Unauthorized
			var resp = map[string]interface{}{"status": false, "message": "Missing auth token"}
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode(resp)
			return
		}
		tk := &models.Token{}

		// fmt.Println("got here", tk)
		// fmt.Println("got here --tokenpart", tokenPart)

		token, err := jwt.ParseWithClaims(authHdr, tk, func(token *jwt.Token) (interface{}, error) {
			// since we only use the one private key to sign the tokens,
			// we also only use its public counter part to verify
			// return "secret", nil
			return []byte("secret"), nil
		})

		fmt.Println("got here afeyeryeyrye", token)
		if err != nil {
			fmt.Println("error")
			fmt.Println(err)
			var resp = map[string]interface{}{"status": false, "message": "Validation failed"}
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode(resp)
			return
		}
		ctx := context.WithValue(r.Context(), "user", token)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

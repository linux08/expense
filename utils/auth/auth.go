package auth

import (
	"encoding/json"
	"fmt"
	"net/http"

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
		fmt.Println(tokenHeader)
		if tokenHeader == "" {
			//Token is missing, returns with error code 403 Unauthorized
			var resp = map[string]interface{}{"status": false, "message": "Missing auth token"}
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode(resp)
			return
		}

		token, err := jwt.Parse(tokenHeader, func(token *jwt.Token) (interface{}, error) {
			// Don't forget to validate the alg is what you expect:
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
			// return _, nil
			fmt.Println("ididi")
			return nil, nil
		})

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			fmt.Println(claims["foo"], claims["nbf"])
		} else {
			fmt.Println(err)
		}

		// splitted := strings.Split(tokenHeader, ".")
		// fmt.Println(splitted)
		// fmt.Println(len(splitted))
		// if len(splitted) == 3 {
		// 	token, error := jwt.Parse(splitted[1], func(token *jwt.Token) (interface{}, error) {
		// 		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		// 			return nil, fmt.Errorf("There was an error")
		// 		}
		// 		return []byte("secret"), nil
		// 	})
		// 	fmt.Println("token")
		// 	fmt.Println("the token")

		// 	fmt.Println(error)
		// 	if error != nil {
		// 		fmt.Println("no error")
		// 		fmt.Println(error.Error)
		// 		var resp = map[string]interface{}{"status": false, "message": error}
		// 		w.WriteHeader(http.StatusForbidden)
		// 		json.NewEncoder(w).Encode(resp)
		// 		return
		// 	}
		// 	if token.Valid {
		// 		fmt.Println("vali token")
		// 		//Everything went well, proceed with the request and set the caller to the user retrieved from the parsed token
		// 		fmt.Sprintf("User %", token.Claims) //Useful for monitoring
		// 		ctx := context.WithValue(r.Context(), "user", token.Claims)
		// 		r = r.WithContext(ctx)
		// 		next.ServeHTTP(w, r)
		// 	} else {
		// 		fmt.Println("others")
		// 		var resp = map[string]interface{}{"status": false, "message": "Invalid/Malformsssed auth token"}
		// 		w.WriteHeader(http.StatusForbidden)
		// 		json.NewEncoder(w).Encode(resp)
		// 	}
		// }
		// fmt.Println("ah ah")
		// var resp = map[string]interface{}{"status": false, "message": "Invalid/Malformed auth token"}
		// w.WriteHeader(http.StatusForbidden)
		// json.NewEncoder(w).Encode(resp)

	})
}

// }
// if len(splitted) != 2 {
// 	var resp = map[string]interface{}{"status": false, "message": "Invalid/Malformed auth token"}
// 	w.WriteHeader(http.StatusForbidden)
// 	json.NewEncoder(w).Encode(resp)
// 	return
// }

// tokenPart := splitted[1] //Grab the token part, what we are truly interested in
// tk := &models.Token{}

// token, err := jwt.ParseWithClaims(tokenPart, tk, func(token *jwt.Token) (interface{}, error) {
// 	return []byte(os.Getenv("token_password")), nil
// })
// fmt.Println("the error")
// fmt.Println(err)
// if err != nil { //Malformed token, returns with http code 403 as usual
// 	var resp = map[string]interface{}{"status": false, "message": "Malformed authentication token"}
// 	w.WriteHeader(http.StatusForbidden)
// 	json.NewEncoder(w).Encode(resp)
// 	return
// }

// if !token.Valid { //Token is invalid, maybe not signed on this server
// 	var resp = map[string]interface{}{"status": false, "message": "Token is not valid"}
// 	w.WriteHeader(http.StatusForbidden)
// 	json.NewEncoder(w).Encode(resp)
// 	return
// }

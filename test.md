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

		authHdr := r.Header.Get("Authorization") //Grab the token from the header
		splitted := strings.Split(authHdr, ".")

		if authHdr == "" {
			//Token is missing, returns with error code 403 Unauthorized
			var resp = map[string]interface{}{"status": false, "message": "Missing auth token"}
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode(resp)
			return
		}
		fmt.Println("toke", splitted)

		tokenPart := splitted[1] //Grab the token part, what we are truly interested in
		fmt.Println("got here")
		tk := &models.Token{}

		fmt.Println("got here", tk)
		fmt.Println("got here --tokenpart", tokenPart)
		//
		//proceed in the middleware chain!

		// splitted := strings.Split(authHdr, ".")
		// fmt.Println(splitted)
		// fmt.Println(len(splitted))
		// if len(splitted) == 3 {
		token, error := jwt.Parse(splitted[2], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("There was an error")
			}
			return []byte("secret"), nil
		})
		fmt.Println("token")
		fmt.Println("the token")

		if error != nil { //Malformed token, returns with http code 403 as usual
			fmt.Println("error", error)
			// fmt.Println("the token")
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
		fmt.Println("got hehhr")
		//Everything went well, proceed with the request and set the caller to the user retrieved from the parsed token
		fmt.Sprintf("User %", tk.UserID) //Useful for monitoring
		ctx := context.WithValue(r.Context(), "user", tk.UserID)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r) //proceed in the middleware chain!
	})
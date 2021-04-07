func testToken(res http.ResponseWriter, req * http.Request){
	expireToken := time.Now().Add(time.Minute * 30)
	expireCookie := time.Now().Add(time.Minute * 30)
	
	claims  := { ... }

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, _ := token.SignedString([]byte("secret"))
}
//json web tokens to generate a session token on the server-side.
//so we can then store and use this token to validate the user and enforce our access controll.
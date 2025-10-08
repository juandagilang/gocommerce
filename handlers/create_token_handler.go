package handlers

import "github.com/dgrijalva/jwt-go"

var jwtKey = []byte("a038fd667db4d281564f9729cf2e86c4972f48f81df0219df124c71143123ddc")

func CreateToken(userID uint) (string, error) {
	claims := jwt.MapClaims{}
	claims["user_id"] = userID
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

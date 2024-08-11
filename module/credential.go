package main

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

var jwtKey = []byte("my_secret_key")

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func main() {
	token := generate_token()
	read_token(token)
	fmt.Println(token)
}

func generate_token() string {
	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &Claims{
		Username: "Lucas",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		fmt.Println("Could not create token")
	}
	return tokenString
}

func read_token(tokenString string) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil || !token.Valid {
		fmt.Println("Invalid token")
		return
	}

	fmt.Println("Welcome ", claims.Username)
}

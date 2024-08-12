package credential

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
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
	token := generate_token("Lucas")
	read_token(token)
	fmt.Println(token)
}

func generate_token(Username string) string {
	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &Claims{
		Username: Username,
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

func Send_token(c *gin.Context) {
	username, password, ok := c.Request.BasicAuth()
	cred := generate_token(username)
	fmt.Println(cred)
	fmt.Println(username, password, ok)
	c.JSON(http.StatusOK, gin.H{"token": cred})
}

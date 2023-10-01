package auth

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt"
	"github.com/jelinden/content-service/app/db"
	"github.com/jelinden/content-service/app/domain"
)

func init() {
	secretKey = os.Getenv("JWT_KEY")
}

type JwtToken struct {
	Token string `json:"token"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	var user domain.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		json.NewEncoder(w).Encode(Exception{Message: err.Error()})
		return
	}
	if validateCredentials(user) {
		json.NewEncoder(w).Encode(JwtToken{Token: signedTokenString(domain.User{Username: user.Username})})
		return
	}
	json.NewEncoder(w).Encode(Exception{Message: "invalid credentials"})
}

func validateCredentials(user domain.User) bool {
	return db.GetUser(user.Username).Password == user.Password
}

func signedTokenString(user domain.User) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS384, jwt.MapClaims{
		"username": user.Username,
	})
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		log.Println(err)
	}
	return tokenString
}

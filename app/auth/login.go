package auth

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"github.com/jelinden/content-service/app/db"
	"github.com/jelinden/content-service/app/domain"
	"github.com/jelinden/content-service/app/util"
)

func init() {
	secretKey = os.Getenv("JWT_KEY")
}

func Login(w http.ResponseWriter, r *http.Request) {
	var user domain.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Exception{Message: err.Error()})
		return
	}
	if validateCredentials(user) {

		authCookie := http.Cookie{
			Name:     "content-service",
			Path:     "/",
			MaxAge:   3600,
			HttpOnly: true,
			Secure:   true,
			SameSite: http.SameSiteLaxMode,
			Value:    signedTokenString(domain.User{Username: user.Username}),
		}
		http.SetCookie(w, &authCookie)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode("OK")
		return
	}
	w.WriteHeader(http.StatusForbidden)
	json.NewEncoder(w).Encode(Exception{Message: "invalid credentials"})
}

func validateCredentials(user domain.User) bool {
	return util.CheckPasswordHash(user.Password, db.GetUser(user.Username).Password)
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

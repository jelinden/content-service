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
	var loggingUser domain.User
	err := json.NewDecoder(r.Body).Decode(&loggingUser)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Exception{Message: err.Error()})
		return
	}
	user := db.GetUser(loggingUser.Username)
	if validateCredentials(loggingUser, user) {

		authCookie := http.Cookie{
			Name:     "content-service",
			Path:     "/",
			MaxAge:   3600,
			HttpOnly: true,
			Secure:   true,
			SameSite: http.SameSiteLaxMode,
			Value:    signedTokenString(domain.User{ID: loggingUser.ID, Username: loggingUser.Username, ApiToken: user.ApiToken}),
		}
		http.SetCookie(w, &authCookie)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode("OK")
		return
	}
	w.WriteHeader(http.StatusForbidden)
	json.NewEncoder(w).Encode(Exception{Message: "invalid credentials"})
}

func validateCredentials(logginUser domain.User, user domain.User) bool {
	return util.CheckPasswordHash(logginUser.Password, user.Password)
}

func signedTokenString(user domain.User) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS384, jwt.MapClaims{
		"id":       user.ID,
		"username": user.Username,
		"apiToken": user.ApiToken,
	})
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		log.Println(err)
	}
	return tokenString
}

func Logout(w http.ResponseWriter, r *http.Request) {
	authCookie := http.Cookie{
		Name:     "content-service",
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
		Value:    "",
	}
	http.SetCookie(w, &authCookie)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("OK")
}

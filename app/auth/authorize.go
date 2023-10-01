package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/context"
	"github.com/julienschmidt/httprouter"
)

var secretKey string

func init() {
	secretKey = os.Getenv("JWT_KEY")
}

func AuthorizeMiddleware(next http.HandlerFunc) httprouter.Handle {
	return httprouter.Handle(func(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
		authCookie, err := req.Cookie("content-service")
		if err != nil {
			json.NewEncoder(w).Encode(Exception{Message: "Invalid Authorization token"})
			return
		}
		authorization := authCookie.Value
		if authorization != "" {
			bearerToken := strings.Split(authorization, " ")
			if len(bearerToken) == 2 {
				token, err := parseBearerToken(bearerToken[1])
				if err != nil {
					json.NewEncoder(w).Encode(Exception{Message: err.Error()})
					return
				}
				if token.Valid {
					context.Set(req, "decoded", token.Claims)
					next(w, req)
					return
				}
				json.NewEncoder(w).Encode(Exception{Message: "Invalid Authorization token"})
				return
			}
		}
		json.NewEncoder(w).Encode(Exception{Message: "Authorization is required"})
	})
}

func parseBearerToken(bearerToken string) (*jwt.Token, error) {
	return jwt.Parse(bearerToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("there was an error")
		}
		return []byte(secretKey), nil
	})
}

type Exception struct {
	Message string `json:"message"`
}

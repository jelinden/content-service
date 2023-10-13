package routes

import (
	"encoding/json"
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/gorilla/context"
	"github.com/jelinden/content-service/app/db"
	"github.com/jelinden/content-service/app/domain"
	"github.com/mitchellh/mapstructure"
)

func AddSpace(w http.ResponseWriter, req *http.Request) {
	var space domain.Space
	err := json.NewDecoder(req.Body).Decode(&space)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user := getUserContext(req)
	newSpace := db.AddSpace(user, space.Name)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(newSpace)
}

func GetSpacesWithUserID(w http.ResponseWriter, req *http.Request) {
	user := getUserContext(req)
	spaces, err := db.GetSpacesWithUserID(user.ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("ERROR"))
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(spaces)
}

func getUserContext(req *http.Request) domain.User {
	decoded := context.Get(req, "decoded")
	var user domain.User
	mapstructure.Decode(decoded.(jwt.MapClaims), &user)
	return user
}

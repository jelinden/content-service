package routes

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/context"
	"github.com/jelinden/content-service/app/db"
	"github.com/jelinden/content-service/app/domain"
	"github.com/julienschmidt/httprouter"
	"github.com/mitchellh/mapstructure"
)

func AddSpace(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
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

func RemoveSpace(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	spaceID, err := strconv.ParseInt(ps.ByName("id"), 10, 64)
	if err != nil {
		log.Println(err)
	} else {
		newSpace := db.RemoveSpace(spaceID)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(newSpace)
		return
	}
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte("Bed request"))
}

func GetSpacesWithUserID(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
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

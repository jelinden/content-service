package routes

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/jelinden/content-service/app/db"
	"github.com/jelinden/content-service/app/domain"
	"github.com/julienschmidt/httprouter"
)

func AddContent(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	var content domain.Content
	err := json.NewDecoder(req.Body).Decode(&content)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	newContent := db.AddContent(content.SpaceID, content.Name, content.Value)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(newContent)
}

func RemoveContent(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	contentID, err := strconv.ParseInt(ps.ByName("id"), 10, 64)
	if err != nil {
		log.Println(err)
	} else {
		newContent := db.RemoveContent(contentID)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(newContent)
		return
	}
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte("Bed request"))
}

func GetContentWithSpaceID(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	spaceID, err := strconv.ParseInt(ps.ByName("spaceID"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	spaces, err := db.GetContentWithSpaceID(spaceID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("ERROR"))
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(spaces)
}

func GetContentWithSpaceIDAndToken(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	spaceID, err := strconv.ParseInt(ps.ByName("spaceID"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	userID, err := strconv.ParseInt(ps.ByName("userID"), 10, 0)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	content, err := db.GetSpaceContentWithUserID(spaceID, userID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("ERROR"))
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(content)
}

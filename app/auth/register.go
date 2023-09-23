package auth

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"

	"github.com/jelinden/content-service/app/domain"
	"github.com/julienschmidt/httprouter"
)

var credentials = make(map[string]string)

func Register(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		handleError(err, "Oops, signup failure", w)
		return
	}
	var user domain.User
	err = json.Unmarshal(body, &user)
	if err != nil {
		handleError(err, "Oops, signup failure", w)
		return
	}
	if credentials[user.Username] != "" {
		handleError(errors.New("username exists"), "Username already exists", w)
		return
	}
	credentials[user.Username] = user.Password
	w.WriteHeader(200)
	w.Write([]byte(`{"status": "OK"}`))
}

func handleError(err error, msg string, w http.ResponseWriter) {
	log.Println(err)
	w.WriteHeader(401)
	e := domain.Error{Error: msg}
	s, _ := json.Marshal(e)
	w.Write(s)
}

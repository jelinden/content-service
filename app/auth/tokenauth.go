package auth

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/jelinden/content-service/app/db"
	"github.com/julienschmidt/httprouter"
)

func TokenAuth(next httprouter.Handle) httprouter.Handle {
	return httprouter.Handle(func(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
		token := req.URL.Query().Get("token")
		user := db.GetUserWithToken(token)
		if user.ApiToken != "" {
			ps = append(ps, httprouter.Param{Key: "userID", Value: strconv.FormatInt(user.ID, 10)})
			next(w, req, ps)
			return
		}
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(Exception{Message: "Invalid token"})
	})
}

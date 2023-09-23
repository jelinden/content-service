package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/golang-jwt/jwt"
	"github.com/gorilla/context"
	"github.com/jelinden/content-service/app/auth"
	"github.com/jelinden/content-service/app/domain"
	"github.com/julienschmidt/httprouter"
	"github.com/mitchellh/mapstructure"
)

const port = 8700

func main() {
	router := httprouter.New()
	router.RedirectFixedPath = true
	router.RedirectTrailingSlash = true
	router.POST("/api/register", auth.Register)
	//router.POST("/api/login", auth.Login)
	router.GET("/api/profile", auth.AuthorizeMiddleware(http.HandlerFunc(protectedEndpoint)))

	router.GET("/", index)
	router.GET("/health", health)

	router.Handler("GET", "/static/*filepath", http.FileServer(http.Dir("public")))
	log.Println("Started a server at port", port)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), router))
}

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	http.ServeFile(w, r, "public/index.html")
}

func health(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func protectedEndpoint(w http.ResponseWriter, req *http.Request) {
	decoded := context.Get(req, "decoded")
	var user domain.User
	mapstructure.Decode(decoded.(jwt.MapClaims), &user)
	json.NewEncoder(w).Encode(user)
}

func currentDirectory() string {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	return filepath.Dir(ex)
}

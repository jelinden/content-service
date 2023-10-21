package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/context"
	"github.com/jelinden/content-service/app/auth"
	"github.com/jelinden/content-service/app/db"
	"github.com/jelinden/content-service/app/domain"
	"github.com/jelinden/content-service/app/routes"
	"github.com/julienschmidt/httprouter"
	"github.com/mitchellh/mapstructure"
)

const port = 8700

func main() {
	db.Init()
	router := httprouter.New()
	router.RedirectFixedPath = true
	router.RedirectTrailingSlash = true
	router.GlobalOPTIONS = http.HandlerFunc(globalOptionsHandler)
	router.POST("/api/register", CorsMiddleware(auth.Register))
	router.POST("/api/login", CorsMiddleware(auth.Login))
	router.POST("/api/logout", CorsMiddleware(auth.Logout))
	router.GET("/api/profile", auth.AuthorizeMiddleware(profile))

	router.POST("/api/space", auth.AuthorizeMiddleware(routes.AddSpace))
	router.DELETE("/api/space/:id", auth.AuthorizeMiddleware(routes.RemoveSpace))
	router.GET("/api/spaces", auth.AuthorizeMiddleware(routes.GetSpacesWithUserID))

	router.POST("/api/content", auth.AuthorizeMiddleware(routes.AddContent))
	router.DELETE("/api/content/:id", auth.AuthorizeMiddleware(routes.RemoveContent))
	router.GET("/api/content/:spaceID", auth.AuthorizeMiddleware(routes.GetContentWithSpaceID))

	router.GET("/", index)
	router.GET("/register", index)
	router.GET("/login", index)
	router.GET("/space", index)
	router.GET("/content/*catchall", index)
	router.GET("/profile", index)
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

func profile(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	decoded := context.Get(req, "decoded")
	log.Println(decoded)
	var user domain.User
	mapstructure.Decode(decoded.(jwt.MapClaims), &user)
	log.Println(user.Username, user.ID)
	json.NewEncoder(w).Encode(user)
}

func globalOptionsHandler(w http.ResponseWriter, req *http.Request) {
	if req.Header.Get("Access-Control-Request-Method") != "" {
		header := w.Header()
		header.Set("Access-Control-Allow-Methods", header.Get("Allow"))
		header.Set("Access-Control-Allow-Headers", req.Header.Get("Access-Control-Request-Headers"))
		header.Set("Access-Control-Allow-Origin", "*")
	}
	w.WriteHeader(http.StatusNoContent)
}

func CorsMiddleware(next http.HandlerFunc) httprouter.Handle {
	return httprouter.Handle(func(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
		w.Header().Add("Access-Control-Allow-Origin", "*")
		next(w, req)
	})
}

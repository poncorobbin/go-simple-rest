package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/joho/godotenv"
	"github.com/poncorobbin/go-simple-rest/pkg/middlewares"

	jwt "github.com/golang-jwt/jwt/v4"
	// gubrak "github.com/novalagung/gubrak/v2"
)

type M map[string]interface{}

var APPLICATION_NAME = "My Simple JWT App"
var LOGIN_EXPIRATION_DURATION = time.Duration(1) * time.Hour
var JWT_SIGNING_METHOD = jwt.SigningMethodHS256
var JWT_SIGNATURE_KEY = []byte("the secret of kalimdor")

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	mux := new(middlewares.CustomMux)
	mux.RegisterMiddleware(middlewares.MiddlewareJWTAuthorization)

	mux.HandleFunc("/login", HandlerLogin)
	mux.HandleFunc("/index", HandlerIndex)

	server := new(http.Server)
	server.Addr = ":8090"
	server.Handler = mux

	fmt.Println("start and running", server.Addr)
	server.ListenAndServe()
}

func HandlerLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Unsupported Method", http.StatusBadRequest)
		return
	}

	userInfo := M{
		"username": "ponco",
		"email":    "ponco@localhost.com",
		"group":    "admin",
	}

	type MyClaims struct {
		jwt.StandardClaims
		Username string `json:"Username"`
		Email    string `json:"Email"`
		Group    string `json:"Group"`
	}

	// create claims/payload
	claims := MyClaims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    APPLICATION_NAME,
			ExpiresAt: time.Now().Add(LOGIN_EXPIRATION_DURATION).Unix(),
		},
		Username: userInfo["username"].(string),
		Email:    userInfo["email"].(string),
		Group:    userInfo["group"].(string),
	}

	// create token
	token := jwt.NewWithClaims(JWT_SIGNING_METHOD, claims)

	// sign token
	signedToken, err := token.SignedString(JWT_SIGNATURE_KEY)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	tokenString, _ := json.Marshal(M{"token": signedToken})
	w.Write([]byte(tokenString))
}

func HandlerIndex(w http.ResponseWriter, r *http.Request) {
	userInfo := r.Context().Value("userInfo").(jwt.MapClaims)
	message := fmt.Sprintf("hello %s (%s)", userInfo["Username"], userInfo["Group"])

	w.Write([]byte(message))
}

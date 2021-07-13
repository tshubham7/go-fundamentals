package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/facebook"
	"github.com/markbates/goth/providers/google"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	key := "Secret-session-key" // Replace with your SESSION_SECRET or similar
	maxAge := 86400 * 30        // 30 days
	isProd := false             // Set to true when serving over https

	store := sessions.NewCookieStore([]byte(key))
	// http.Cookie÷ß
	store.MaxAge(maxAge)
	store.Options.Path = "/"
	store.Options.HttpOnly = true // HttpOnly should always be enabled
	store.Options.Secure = isProd

	gothic.Store = store

	goth.UseProviders(
		facebook.New(os.Getenv("FB_KEY"), os.Getenv("FB_SECRET"), os.Getenv("FB_CALLBACKURL"), "email"),
		google.New(os.Getenv("GOOLE_KEY"), os.Getenv("GOOLE_SECRET"), os.Getenv("GOOLE_CALLBACKURL"), "email", "profile"),
	)

	r := mux.NewRouter()

	r.HandleFunc("/auth/callback", completeauth)
	r.HandleFunc("/auth/{provider}", beginauth)
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"message": "social auth example"}`))
	})

	addr := ":8000"

	srv := &http.Server{
		Handler: r,
		Addr:    addr,
	}
	fmt.Println("listening on", addr)
	log.Fatal(srv.ListenAndServe())
}

func beginauth(w http.ResponseWriter, r *http.Request) {
	gothic.BeginAuthHandler(w, r) //authentication with provider
}

func completeauth(w http.ResponseWriter, r *http.Request) {

	user, err := gothic.CompleteUserAuth(w, r) //get autherised data's (name,id,profile)
	if err != nil {
		w.Write([]byte(`{"error": "unable to handle user data"}`))
		return
	}
	// do whatever you want with the user data

	b, err := json.Marshal(&user)
	if err != nil {
		w.Write([]byte(`{"error": "unable to handle user data"}`))
		return
	}
	w.Write(b)
}

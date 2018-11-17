package main

import (
	"log"
	"net/http"
	"fmt"
	"github.com/alknopfler/alkalarm/config"
	"io/ioutil"
	"time"
	"encoding/base64"
	"math/rand"
	"context"
)

const oauthGoogleUrlAPI = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="

func oauthByGoogleOauth(w http.ResponseWriter, r *http.Request){
	// Create oauthState cookie
	oauthState := generateStateOauthCookie(w)
	u := config.GConfAuth.AuthCodeURL(oauthState)
	http.Redirect(w, r, u, http.StatusTemporaryRedirect)
}

func generateStateOauthCookie(w http.ResponseWriter) string {
	var expiration = time.Now().Add(365 * 24 * time.Hour)

	b := make([]byte, 16)
	rand.Read(b)
	state := base64.URLEncoding.EncodeToString(b)
	cookie := http.Cookie{Name: "oauthstate", Value: state, Expires: expiration}
	http.SetCookie(w, &cookie)

	return state
}


func oauthGoogleCallback(w http.ResponseWriter, r *http.Request) {
	// Read oauthState from Cookie
	oauthState, _ := r.Cookie("oauthstate")

	if r.FormValue("state") != oauthState.Value {
		log.Println("invalid oauth google state")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	data, err := getUserDataFromGoogle(r.FormValue("code"))
	if err != nil {
		log.Println(err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	// GetOrCreate User in your db.
	// Redirect or response with a token.
	// More code .....
	fmt.Fprintf(w, "UserInfo: %s\n", data)
}

func getUserDataFromGoogle(code string) ([]byte, error) {
	// Use code to get token and get user info from Google.

	token, err := config.GConfAuth.Exchange(context.Background(), code)
	if err != nil {
		return nil, fmt.Errorf("code exchange wrong: %s", err.Error())
	}
	response, err := http.Get(oauthGoogleUrlAPI + token.AccessToken)
	if err != nil {
		return nil, fmt.Errorf("failed getting user info: %s", err.Error())
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed read response: %s", err.Error())
	}
	return contents, nil
}


func new() http.Handler {
	mux := http.NewServeMux()
	// Root
	mux.Handle("/",  http.FileServer(http.Dir("./")))
	// OauthGoogle
	mux.HandleFunc("/auth", oauthByGoogleOauth)
	mux.HandleFunc("/callback", oauthGoogleCallback)

	return mux
}

func main() {
	server := &http.Server{
		Addr: fmt.Sprintf(":80"),
		Handler: new(),
	}

	log.Printf("Starting HTTP Server. Listening at %q", server.Addr)
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.Printf("%v", err)
	} else {
		log.Println("Server closed!")
	}
}



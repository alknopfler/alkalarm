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
	"github.com/gorilla/mux"
	"encoding/json"
	"strings"
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
		http.Redirect(w, r, "/index.html", http.StatusTemporaryRedirect)
		return
	}

	data, err := getUserDataFromGoogle(r.FormValue("code"))
	if err != nil {
		log.Println(err.Error())
		http.Redirect(w, r, "/index.html", http.StatusTemporaryRedirect)
		return
	}
	var gmodel config.GoogleUser
	json.Unmarshal(data,&gmodel)

	if strings.Contains(config.LIST_ACCESS, gmodel.Email){
		http.Redirect(w,r,"/static/index.html",http.StatusFound)
	}else{
		http.Redirect(w,r,"/static/fail.html",http.StatusNotFound)
	}
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

func MyHandler (w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[1:]
	log.Println(path)
	var data []byte
	var err error

	data, err = ioutil.ReadFile("./index.html")

	if err == nil {
		w.Write(data)
	} else {
		w.WriteHeader(404)
		w.Write([]byte("404 Something went wrong - " + http.StatusText(404)))
	}
}
func main() {
	r := mux.NewRouter()
	// Root
	r.HandleFunc("/", MyHandler)
	r.PathPrefix("/static/").Handler(
		http.StripPrefix("/static/", http.FileServer(http.Dir("/opt/alkalarm/webinterface/static/"))),
	)
	//mux.Handle("/static/",  http.FileServer(http.Dir("/opt/alkalarm/webinterface/static")))
	// OauthGoogle
	r.HandleFunc("/auth", oauthByGoogleOauth)
	r.HandleFunc("/callback", oauthGoogleCallback)
	err := http.ListenAndServe(":80",r)
	if err != nil {
		log.Println("Error listening api server...")
	}
}



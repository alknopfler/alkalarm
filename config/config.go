package config

import (
	"io/ioutil"
	"log"
	"fmt"
	"os"
	"encoding/json"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

const (
	//GENERIC
	PROJECT_PATH	="/opt/alkalarm/"
	PYGPIO		="_433.py"
	SERVER_API_PORT =":8080"
	//DATABASE

	//CONTROL
	STATE_FULL	="full"
	STATE_PART	="partial"
	STATE_INAC	="inactive"
	STATE_SOS	="sos"
)

var (
	DB_NAME = PROJECT_PATH+"data.db"
	DB_TYPE	= "sqlite3"

	FROM = readFromFile(".userSMTP")
	SMTP_SERVER = "smtp.gmail.com"
	SMTP_PORT = "587"
	SMTP_PASS = readFromFile(".passSMTP")
	LIST_ACCESS = readFromFile(".listACCESS")
	Cred = GetOauthCred()
	GConfAuth = &oauth2.Config{
			ClientID:     Cred.Cid,
			ClientSecret: Cred.Csecret,
			RedirectURL:  "http://alknopfler.ddns.net/callback",
			Scopes: []string{
							"https://www.googleapis.com/auth/userinfo.profile",
							"https://www.googleapis.com/auth/userinfo.email", // You have to select your own scope from here -> https://developers.google.com/identity/protocols/googlescopes#google_sign-in
							 },
			Endpoint: google.Endpoint,
			}

)
type ArraySensor struct{
	Data []Sensor `json:"Data"`
}
//Sensor struct to define the object
type Sensor struct {
	Code string  `json:"Code"`
	TypeOf string `json:"TypeOf"`
	Zone string `json:"Zone"`
}

//Mailer struct to define the object
type Mailer struct {
	Receptor string  `json:"Receptor"`
}

type ArrayControl struct{
	Data []Control `json:"Data"`
}
//Control struct to define the object
type Control struct {
	Code string  `json:"Code"`
	Description string `json:"Description"`
	TypeOf string `json:"TypeOf"`
}

type Alarm struct{
	Date string  `json:"Date"`
	Sensor string `json:"Sensor"`
}

type GlobalState struct{
	Id string `json:"Id"`
	GState string `json:"State"`
}

type Credentials struct {
	Cid string `json:"cid"`
	Csecret string `json:"csecret"`
}

type GoogleUser struct {
	ID string `json:"id"`
	Email string `json:"email"`
	VerifiedEmail bool `json:"verified_email"`
	Name string `json:"name"`
	GivenName string `json:"given_name"`
	FamilyName string `json:"family_name"`
	Link string `json:"link"`
	Picture string `json:"picture"`
	Gender string `json:"gender"`
	Locale string `json:"locale"`
}

func GetOauthCred() Credentials{
	var c Credentials
	file, err := ioutil.ReadFile(PROJECT_PATH+"creds.json")
	if err != nil {
		fmt.Printf("File error: %v\n", err)
		os.Exit(1)
	}
	json.Unmarshal(file, &c)
	return c
}

func readFromFile(file string)string{
	b, err := ioutil.ReadFile(PROJECT_PATH+file) // just pass the file name
	if err != nil {
		log.Print(err)
	}
	return string(b)
}
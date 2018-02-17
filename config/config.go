package config

import (
	"io/ioutil"
	"log"
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
	WEBACCESS_PASS = readFromFile(".passACCESS")

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


func readFromFile(file string)string{
	b, err := ioutil.ReadFile(PROJECT_PATH+file) // just pass the file name
	if err != nil {
		log.Print(err)
	}
	return string(b)
}
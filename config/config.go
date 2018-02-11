package config

import (
	"io/ioutil"
	"fmt"
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

	FROM = "alknopfler@gmail.com"
	SMTP_USER = "alknopfler@gmail.com"
	LIST_TO_MAIL = []string{"alknopfler@gmail.com","begoclavero@gmail.com"}
	SMTP_SERVER = "smtp.gmail.com"
	SMTP_PORT = "587"
	SMTP_PASS = readPassFromFile()

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


func readPassFromFile()string{
	b, err := ioutil.ReadFile(PROJECT_PATH+".passSMTP") // just pass the file name
	if err != nil {
		fmt.Print(err)
	}
	return string(b)
}
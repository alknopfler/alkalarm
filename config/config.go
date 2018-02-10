package config

import (
	"io/ioutil"
	"fmt"
)

const (
	//GENERIC
	PROJECT_PATH	="/opt/alkalarm/"
	PYGPIO		="_433.py"
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

//Sensor struct to define the object
type Sensor struct {
	Code string
	TypeOf string
	Zone string
}

func readPassFromFile()string{
	b, err := ioutil.ReadFile(PROJECT_PATH+".passSMTP") // just pass the file name
	if err != nil {
		fmt.Print(err)
	}
	return string(b)
}
package config

import "os"

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
	DB_NAME = os.Getenv("HOME")+"/data.db"
	DB_TYPE	="sqlite3"
)

//Sensor struct to define the object
type Sensor struct {
	Code string
	TypeOf string
	Zone string
}
//Mailer struct to define the object
type Mailer struct{
	Emisor string
	Receptor string
	Subject string
	Text string
	Smtp_address string
	Smtp_port string
	Smtp_user string
	Smtp_pass string
	Smtp_security string
}


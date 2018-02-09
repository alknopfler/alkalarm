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
package database

import (
	"database/sql"
	"github.com/alknopfler/alkalarm/config"
	"fmt"
)

func CreateSchemas(){
	db, err := sql.Open(config.DB_TYPE, config.DB_PATH)
	if err!=nil{
		fmt.Println("Error opening Database: ",err)
	}
	task, _:= db.Prepare(SENSORS_TABLE)
	task.Exec()

	task, _= db.Prepare(MAILER_TABLE)
	task.Exec()

	task, _= db.Prepare(ALARM_HISTORY_TABLE)
	task.Exec()

	task, _= db.Prepare(CONTROL_CODES_TABLE)
	task.Exec()

	task, _= db.Prepare(GLOBAL_STATE)
	task.Exec()

	db.Close()
}
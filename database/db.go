package database

import (
	"database/sql"
	"github.com/alknopfler/alkalarm/config"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

func CreateSchemas(){
	db, err := sql.Open(config.DB_TYPE, config.DB_PATH)
	if err!=nil{
		fmt.Println("Error opening Database: ",err)
	}
	task, _:= db.Prepare(SENSORS_TABLE)
	task.Exec()
	fmt.Println("primera")
	task, _= db.Prepare(MAILER_TABLE)
	task.Exec()
	fmt.Println("segunda")

	task, _= db.Prepare(ALARM_HISTORY_TABLE)
	task.Exec()
	fmt.Println("tercera")

	task, _= db.Prepare(CONTROL_CODES_TABLE)
	task.Exec()
	fmt.Println("cuarta")

	task, _= db.Prepare(GLOBAL_STATE)
	task.Exec()
	fmt.Println("quinta")

	db.Close()
}
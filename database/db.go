package database

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

func CreateSchemas(){
	db, err := sql.Open("sqlite3", "./alkdb.db")
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
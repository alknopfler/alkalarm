package database

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"github.com/alknopfler/alkalarm/config"
)
func InitDB(filepath string) *sql.DB {
	db, err := sql.Open(config.DB_TYPE, config.DB_NAME)
	if err != nil { panic(err) }
	if db == nil { panic("db nil") }
	return db
}

func CreateSchemas(db *sql.DB){

	list:=[]string{SENSORS_TABLE,MAILER_TABLE,ALARM_HISTORY_TABLE,CONTROL_CODES_TABLE,GLOBAL_STATE_TABLE}
	for _,val := range list{
		db.Exec(val)
	}
}

func OperateWithItem(db *sql.DB, operation string, values ...interface{}){
	line, err := db.Prepare(operation)
	if err != nil {
		fmt.Println("Error preparing the DB for the operation: ",err)
		return
	}
	defer line.Close()
	_, err2 := line.Exec(values...)
	if err2 != nil{
		fmt.Println("Error executing the operation in DB: ",err2)
		return
	}
	fmt.Println("Operation '%s' in DB successfully done...", operation)
}
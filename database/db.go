package database

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"

	"github.com/alknopfler/alkalarm/config"
)
func InitDB() (*sql.DB,error) {
	db, err := sql.Open(config.DB_TYPE, config.DB_NAME)
	if err != nil { return nil,err }
	return db,nil
}

func CreateSchemas(db *sql.DB) error {
	var err error
	list:=[]string{SENSORS_TABLE,MAILER_TABLE,ALARM_HISTORY_TABLE,CONTROL_CODES_TABLE,GLOBAL_STATE_TABLE}
	for _,val := range list{
		_,err=db.Exec(val)
	}
	return err
}

func OperateWithItem(db *sql.DB, operation string, values ...interface{})error{
	line, err := db.Prepare(operation)
	if err != nil {
		fmt.Println("Error preparing the DB for the operation: ",err)
		return err
	}
	defer line.Close()
	_, err2 := line.Exec(values...)
	if err2 != nil{
		fmt.Println("Error executing the operation in DB: ",err2)
		return err2
	}
	return nil
}
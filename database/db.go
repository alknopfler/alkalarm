package database

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"

	"github.com/alknopfler/alkalarm/config"
)
func InitDB(path ...string) (*sql.DB,error) {
	if path == nil{
		db, err := sql.Open(config.DB_TYPE, config.DB_NAME)
		if err != nil { return nil,err }
		return db,nil
	}else{ //for testing cases
		db, err := sql.Open(config.DB_TYPE, path[0])
		if err != nil { return nil,err }
		return db,nil
	}

}

func CreateSchemas(db *sql.DB) error {
	var err error
	list:=[]string{SENSORS_TABLE,MAILER_TABLE,ALARM_HISTORY_TABLE,CONTROL_CODES_TABLE,GLOBAL_STATE_TABLE}
	for _,val := range list{
		_,err=db.Exec(val)
	}
	return err
}

func Operate(db *sql.DB, operation string, values ...interface{})error{
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
/*
func QuerySensors(db *sql.DB, query string) ([]string,error){
	var result []string
	rows, err := db.Query(query)
	if err != nil { return result,err }
	defer rows.Close()


	for rows.Next() {
		item := string{}
		err2 := rows.Scan(&item.Code, &item.TypeOf, &item.Zone)
		if err2 != nil { return nil,err }
		result = append(result, item)
	}
	return result, nil
}*/


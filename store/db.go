package store

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"

	cfg "github.com/alknopfler/alkalarm/config"
	"github.com/alknopfler/alkalarm/sensors"
)
func InitDB(path ...string) (*sql.DB,error) {
	if path == nil{
		db, err := sql.Open(cfg.DB_TYPE, cfg.DB_NAME)
		if err != nil { return nil,err }
		return db,nil
	}else{ //for testing cases
		db, err := sql.Open(cfg.DB_TYPE, path[0])
		if err != nil { return nil,err }
		return db,nil
	}

}

func CreateSchemas(db *sql.DB) error {
	var err error
	list:=[]string{cfg.SENSORS_TABLE,cfg.MAILER_TABLE,cfg.ALARM_HISTORY_TABLE,cfg.CONTROL_CODES_TABLE,cfg.GLOBAL_STATE_TABLE}
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

func QuerySensors(db *sql.DB, query string) ([]sensors.Sensor,error){
	var result []sensors.Sensor
	rows, err := db.Query(query)
	if err != nil { return result,err }
	defer rows.Close()


	for rows.Next() {
		item := sensors.Sensor{}
		err2 := rows.Scan(&item.Code, &item.TypeOf, &item.Zone)
		if err2 != nil { return nil,err }
		result = append(result, item)
	}
	return result, nil
}


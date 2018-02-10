package sensors

import (
	cfg "github.com/alknopfler/alkalarm/config"
	"fmt"
	"github.com/alknopfler/alkalarm/database"
)


func RegisterSensor(data cfg.Sensor) error{
	db,err := database.InitDB()
	if err != nil {
		fmt.Println("Error initiating DB in Register Sensor")
		return err
	}
	defer db.Close()

	err=database.Operate(db,cfg.SENSOR_INSERT,data.Code,data.TypeOf,data.Zone)
	if err!=nil{
		fmt.Println("Error inserting sensor in db")
		return err
	}
	fmt.Println("Success...Sensor registered successfully")
	return nil
}

func UnregisterSensor(code string) error{
	db,err := database.InitDB()
	if err != nil {
		fmt.Println("Error initiating DB in Register Sensor")
		return err
	}
	defer db.Close()

	err=database.Operate(db,cfg.SENSOR_DELETE,code)
	if err!=nil{
		fmt.Println("Error inserting sensor in db")
		return err
	}
	fmt.Println("Success...Sensor registered successfully")
	return nil
}


func QuerySensorsAll() ([]cfg.Sensor,error){
	var result []cfg.Sensor
	db,err := database.InitDB()
	if err != nil {
		fmt.Println("Error initiating DB in Query Sensor")
		return result,err
	}
	defer db.Close()
	rows, err := db.Query(cfg.SENSOR_QUERY_ALL)
	if err != nil { return result,err }
	defer rows.Close()

	for rows.Next() {
		item := cfg.Sensor{}
		err2 := rows.Scan(&item.Code, &item.TypeOf, &item.Zone)
		if err2 != nil { return nil,err }
		result = append(result, item)
	}
	return result, nil
}

func QuerySensors(code string) (cfg.Sensor,error){
	var result cfg.Sensor
	db,err := database.InitDB()
	if err != nil {
		fmt.Println("Error initiating DB in Query Sensor")
		return result,err
	}
	defer db.Close()
	rows, err := db.Query(cfg.SENSOR_QUERY_CODE,code)
	if err != nil { return result,err }
	defer rows.Close()

	if rows.Next() {
		err2 := rows.Scan(&result.Code, &result.TypeOf, &result.Zone)
		if err2 != nil { return result,err }
	}
	return result, nil
}

func SensorExists(code string) bool{
	db,err := database.InitDB()
	if err != nil {
		fmt.Println("Error initiating DB in Sensor Exists")
		return false
	}
	defer db.Close()
	rows, err := db.Query(cfg.SENSOR_QUERY_CODE,code)
	defer rows.Close()
	if rows.Next(){
		return true
	}
	return false
}

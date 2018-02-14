package sensors

import (
	cfg "github.com/alknopfler/alkalarm/config"
	"log"
	"github.com/alknopfler/alkalarm/database"
)


func Register(data cfg.Sensor) error{
	db,err := database.InitDB()
	if err != nil {
		log.Println("Error initiating DB in Register Sensor")
		return err
	}
	defer db.Close()

	err=database.Operate(db,cfg.SENSOR_INSERT,data.Code,data.TypeOf,data.Zone)
	if err!=nil{
		log.Println("Error inserting sensor in db")
		return err
	}
	log.Println("Success...Sensor registered successfully")
	return nil
}

func Unregister(code string) error{
	db,err := database.InitDB()
	if err != nil {
		log.Println("Error initiating DB in Register Sensor")
		return err
	}
	defer db.Close()

	err=database.Operate(db,cfg.SENSOR_DELETE,code)
	if err!=nil{
		log.Println("Error inserting sensor in db")
		return err
	}
	log.Println("Success...Sensor unregistered successfully")
	return nil
}


func QueryAll() ([]cfg.Sensor,error){
	var result []cfg.Sensor
	db,err := database.InitDB()
	if err != nil {
		log.Println("Error initiating DB in Query Sensor")
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

func Query(code string) (cfg.Sensor,error){
	var result cfg.Sensor
	db,err := database.InitDB()
	if err != nil {
		log.Println("Error initiating DB in Query Sensor")
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

func Exists(code string) bool{
	db,err := database.InitDB()
	if err != nil {
		log.Println("Error initiating DB in Sensor Exists")
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

func IsPartial(code string) bool{
	s,err:=Query(code)
	if err!=nil{
		return false
	}
	if s.TypeOf == "aperture"{
		return true
	}
	return false
}
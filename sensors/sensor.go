package sensors

import (
	cfg "github.com/alknopfler/alkalarm/config"
	"fmt"
	"github.com/alknopfler/alkalarm/database"
)

func RegisterSensor(data cfg.Sensor) error{
	//First Time to execute needs create database and scheme
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

func UnregisterSensor(data cfg.Sensor) error{
	//First Time to execute needs create database and scheme
	db,err := database.InitDB()
	if err != nil {
		fmt.Println("Error initiating DB in Register Sensor")
		return err
	}
	defer db.Close()

	err=database.Operate(db,cfg.SENSOR_DELETE,data.Code)
	if err!=nil{
		fmt.Println("Error inserting sensor in db")
		return err
	}
	fmt.Println("Success...Sensor registered successfully")
	return nil
}
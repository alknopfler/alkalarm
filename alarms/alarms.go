package alarms

import (
	cfg "github.com/alknopfler/alkalarm/config"
	"fmt"
	"github.com/alknopfler/alkalarm/database"
)


func RegisterAlarm(data cfg.Alarm) error{
	db,err := database.InitDB()
	if err != nil {
		fmt.Println("Error initiating DB in Register Alarm")
		return err
	}
	defer db.Close()

	err=database.Operate(db,cfg.ALARM_INSERT,data.Date,data.Sensor)
	if err!=nil{
		fmt.Println("Error inserting Alarm in db")
		return err
	}
	fmt.Println("Success...Alarm registered successfully")
	return nil
}

func UnregisterAlarm() error{
	db,err := database.InitDB()
	if err != nil {
		fmt.Println("Error initiating DB in unregister Alarm")
		return err
	}
	defer db.Close()

	err=database.Operate(db,cfg.ALARM_DELETE)
	if err!=nil{
		fmt.Println("Error inserting alarm in db")
		return err
	}
	fmt.Println("Success...alarm registered successfully")
	return nil
}


func QueryAlarmAll() ([]cfg.Alarm,error){
	var result []cfg.Alarm
	db,err := database.InitDB()
	if err != nil {
		fmt.Println("Error initiating DB in Query alarm")
		return result,err
	}
	defer db.Close()
	rows, err := db.Query(cfg.ALARM_QUERY_ALL)
	if err != nil { return result,err }
	defer rows.Close()

	for rows.Next() {
		item := cfg.Alarm{}
		err2 := rows.Scan(&item.Date, &item.Sensor)
		if err2 != nil { return nil,err }
		result = append(result, item)
	}
	return result, nil
}


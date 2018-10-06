package alarms

import (
	cfg "github.com/alknopfler/alkalarm/config"
	"log"
	"github.com/alknopfler/alkalarm/database"
	"sync"
)


func Register(data cfg.Alarm, wg *sync.WaitGroup) error{
	defer wg.Done()

	db,err := database.InitDB()
	if err != nil {
		log.Println("Error initiating DB in Register Alarm")
		return err
	}
	defer db.Close()

	err=database.Operate(db,cfg.ALARM_INSERT,data.Date,data.Sensor)
	if err!=nil{
		log.Println("Error inserting Alarm in db")
		return err
	}
	log.Println("Success...Alarm registered successfully")
	return nil
}

func Unregister() error{
	db,err := database.InitDB()
	if err != nil {
		log.Println("Error initiating DB in unregister Alarm")
		return err
	}
	defer db.Close()

	err=database.Operate(db,cfg.ALARM_DELETE)
	if err!=nil{
		log.Println("Error inserting alarm in db")
		return err
	}
	log.Println("Success...alarm unregistered successfully")
	return nil
}


func QueryAll() ([]cfg.Alarm,error){
	var result []cfg.Alarm
	db,err := database.InitDB()
	if err != nil {
		log.Println("Error initiating DB in Query alarm")
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


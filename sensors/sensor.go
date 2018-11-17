package sensors

import (
	"bufio"
	"errors"
	"fmt"
	cfg "github.com/alknopfler/alkalarm/config"
	"github.com/alknopfler/alkalarm/database"
	"log"
	"os/exec"
	"strings"
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




func discoverCodeSensor() (string, error) {
	cmdName := "python -u " + cfg.PROJECT_PATH + cfg.PYGPIO
	cmdArgs := strings.Fields(cmdName)

	cmd := exec.Command(cmdArgs[0], cmdArgs[1:]...)
	stdout, _ := cmd.StdoutPipe()
	cmd.Start()
	oneByte := make([]byte,0)
	for{
		_, err := stdout.Read(oneByte)
		if err != nil {
			fmt.Printf(err.Error())
			break
		}
		r := bufio.NewReader(stdout)
		line,_, _ := r.ReadLine()
		if string(line)!="" {
			cmd.Process.Kill()
			return string(line),nil  //si ha encontrado  salgo directametne con nil matando previamente
		}

	}
	cmd.Process.Kill()  //si no ha encontrado nada en 10 seg mato y salgo con error
	return "",errors.New("Sensor Not Found after 10 second looking for it...")
}

func ScanSensor()(cfg.Sensor,error){
	fmt.Println("Looking for new sensor...Try to activate manually to detect it...")
	code,err:=discoverCodeSensor()
	var sensor cfg.Sensor
	if err!=nil{
		return cfg.Sensor{},err
	}
	fmt.Println("Sensor detected!!! with code: "+code)

	sensor = cfg.Sensor{code,"XXXX","XXXX"}


	return sensor, nil
}

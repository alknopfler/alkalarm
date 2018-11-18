package control

import (
	cfg "github.com/alknopfler/alkalarm/config"
	"log"
	"github.com/alknopfler/alkalarm/database"
	"fmt"
	"bufio"
	"errors"
	"strings"
	"os/exec"
)


func Register(data cfg.Control) error{
	db,err := database.InitDB()
	if err != nil {
		log.Println("Error initiating DB in Register Control")
		return err
	}
	defer db.Close()

	err=database.Operate(db,cfg.CONTROL_INSERT,data.Code,data.Description,data.TypeOf)
	if err!=nil{
		log.Println("Error inserting control in db")
		return err
	}
	log.Println("Success...Control registered successfully")
	return nil
}

func Unregister(code string) error{
	db,err := database.InitDB()
	if err != nil {
		log.Println("Error initiating DB in Register Control")
		return err
	}
	defer db.Close()

	err=database.Operate(db,cfg.CONTROL_DELETE,code)
	if err!=nil{
		log.Println("Error inserting control in db")
		return err
	}
	log.Println("Success...control unregistered successfully")
	return nil
}


func QueryAll() ([]cfg.Control,error){
	var result []cfg.Control
	db,err := database.InitDB()
	if err != nil {
		log.Println("Error initiating DB in Query control")
		return result,err
	}
	defer db.Close()
	rows, err := db.Query(cfg.CONTROL_QUERY_ALL)
	if err != nil { return result,err }
	defer rows.Close()

	for rows.Next() {
		item := cfg.Control{}
		err2 := rows.Scan(&item.Code, &item.Description, &item.TypeOf)
		if err2 != nil { return nil,err }
		result = append(result, item)
	}
	return result, nil
}

func Query(code string) (cfg.Control,error){
	var result cfg.Control
	db,err := database.InitDB()
	if err != nil {
		log.Println("Error initiating DB in Query Control")
		return result,err
	}
	defer db.Close()
	rows, err := db.Query(cfg.CONTROL_QUERY_CODE,code)
	if err != nil { return result,err }
	defer rows.Close()

	if rows.Next() {
		err2 := rows.Scan(&result.Code, &result.Description, &result.TypeOf)
		if err2 != nil { return result,err }
	}
	return result, nil
}

func QueryTypeOf(code string) string {
	c,_:=Query(code)
	return c.TypeOf
}

func Exists(code string) bool{
	db,err := database.InitDB()
	if err != nil {
		log.Println("Error initiating DB in Control Exists")
		return false
	}
	defer db.Close()
	rows, err := db.Query(cfg.CONTROL_QUERY_CODE,code)
	defer rows.Close()
	if rows.Next(){
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

func ScanControl()(cfg.Control,error){
	fmt.Println("Looking for new control...Try to activate manually to detect it...")
	code,err:=discoverCodeSensor()
	var control cfg.Control
	if err!=nil{
		return cfg.Control{},err
	}
	fmt.Println("Control key detected!!! with code: "+code)

	control = cfg.Control{code,"XXXX","XXXX"}


	return control, nil
}

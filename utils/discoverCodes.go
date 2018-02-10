package main

import (
	"strings"
	"os/exec"
	"fmt"
	"bufio"
	"errors"
	cfg "github.com/alknopfler/alkalarm/config"

	"encoding/json"
)

func discoverCodeSensor() (string, error) {
	//TODO en registrar no puede estar activa la alarma
	cmdName := "python -u " + cfg.PROJECT_PATH + cfg.PYGPIO
	cmdArgs := strings.Fields(cmdName)

	cmd := exec.Command(cmdArgs[0], cmdArgs[1:len(cmdArgs)]...)
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
		if (string(line)!=""){
			cmd.Process.Kill()
			return string(line),nil  //si ha encontrado  salgo directametne con nil matando previamente
		}

	}
	cmd.Process.Kill()  //si no ha encontrado nada en 10 seg mato y salgo con error
	return "",errors.New("Sensor Not Found after 10 second looking for it...")
}

func menu(){
	fmt.Println("***********************************")
	fmt.Println("Utility to discover sensor codes...")
	fmt.Println("The program will stop automatically when it discover the sensor.")
	fmt.Println("Anyway, you could stop using ctrl+c")
	fmt.Println("***********************************")
	var listSensor []cfg.Sensor
	for true{
		fmt.Println("Looking for new sensor...Try to activate manually to detect it...")
		code,err:=discoverCodeSensor()

		if err!=nil{
			fmt.Println("Fail: ",err)
		}else{
			fmt.Println("Code Detected:",code)
			fmt.Print("Enter type of Sensor [presence|aperture|other]: ")
			var typeof string
			fmt.Scanln(&typeof)
			fmt.Print("Enter the zone of your house: ")
			var zone string
			fmt.Scanln(&zone)
			listSensor = append(listSensor,cfg.Sensor{code,typeof,zone})
		}
		fmt.Print("Do you want to continuos with another sensor?:[Y|n] ")
		var answer string
		fmt.Scanln(&answer)
		if answer == "n" || answer == "N"{
			break
		}
	}
	response, _ := json.Marshal(listSensor)
	fmt.Println("The json with the sensors detected is: ")
	fmt.Println(response)
	fmt.Println("Try to register in AlkAlarm using the API setup/sensor and load the body payload with this json.")
}

func main(){
	menu()
}

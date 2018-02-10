package main

import (
	"strings"
	"os/exec"
	"fmt"
	"bufio"
	"errors"
	cfg "github.com/alknopfler/alkalarm/config"

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

func main(){
	code,err:=discoverCodeSensor()
	if err!=nil{
		fmt.Println("Fail: ",err)
	}else{
		fmt.Println("Code: ",code)
	}
}

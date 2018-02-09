package sensors

import (
	"strings"
	"os/exec"
	"fmt"
	"bufio"
	"github.com/alknopfler/alkalarm/config"

	"github.com/alknopfler/alkalarm/database"
)
//Sensor struct to define the object
type Sensor struct {
	Code string
	TypeOf string
	Zone string
}

func RegisterSensor(){
	//TODO en registrar no puede estar activa la alarma
	cmdName := "python -u" + config.PROJECT_PATH + config.PYGPIO
	cmdArgs := strings.Fields(cmdName)

	cmd := exec.Command(cmdArgs[0], cmdArgs[1:len(cmdArgs)]...)
	stdout, _ := cmd.StdoutPipe()
	cmd.Start()
	oneByte := make([]byte,0)
	now:=1
	ending:=10
	fmt.Println("before for")
	for (now < ending){
		_, err := stdout.Read(oneByte)
		if err != nil {
			fmt.Printf(err.Error())
			break
		}
		r := bufio.NewReader(stdout)
		line,_, _ := r.ReadLine()
		fmt.Println("la linea es",string(line))
		if string(line) != ""{
			err:=handlerEvent(string(line))
			if err!=nil{
				fmt.Println("Error registering the Sensor")
				return
			}
			break
		}

	}
	cmd.Process.Kill()
	fmt.Println("Success registering the sensor")
}

func UnregisterSensor(){

}

func handlerEvent(evento string) error{
	fmt.Println("Sensor detected: " ,evento)
	db,err:=database.InitDB()
	if err!= nil {
		fmt.Println("Error opening the DB to register sensor: ", err)
		return err
	}
	defer db.Close()


	return nil
}

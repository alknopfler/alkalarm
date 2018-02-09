package sensors

import (
	"strings"
	"os/exec"
	"fmt"
	"bufio"
	"github.com/alknopfler/alkalarm/config"

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
	for (now < ending){
		_, err := stdout.Read(oneByte)
		if err != nil {
			fmt.Printf(err.Error())
			break
		}
		r := bufio.NewReader(stdout)
		line,_, _ := r.ReadLine()
		if string(line) != ""{
			handlerEvent(string(line))
			break
		}

	}
	cmd.Process.Kill()
}

func UnregisterSensor(){

}

func handlerEvent(evento string){

	fmt.Println("Sensor detected: " ,evento)

}


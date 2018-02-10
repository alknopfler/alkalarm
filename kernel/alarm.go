package alarms

import (
	cfg "github.com/alknopfler/alkalarm/config"
	"fmt"
	"strings"
	"os/exec"
	"bufio"
)

func handlerEvent(evento string){

	fmt.Println("Sensor detected: " ,evento)

}

func listenEvents(){
	cmdName := "python -u " + cfg.PROJECT_PATH + cfg.PYGPIO
	cmdArgs := strings.Fields(cmdName)

	cmd := exec.Command(cmdArgs[0], cmdArgs[1:len(cmdArgs)]...)
	stdout, _ := cmd.StdoutPipe()
	cmd.Start()
	oneByte := make([]byte,0)
	for {
		_, err := stdout.Read(oneByte)
		if err != nil {
			fmt.Printf(err.Error())
			break
		}
		r := bufio.NewReader(stdout)
		line,_, _ := r.ReadLine()
		handlerEvent(string(line))

	}

	cmd.Wait()
}

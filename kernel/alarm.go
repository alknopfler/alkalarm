package kernel

import (
	cfg "github.com/alknopfler/alkalarm/config"
	"fmt"
	"strings"
	"os/exec"
	"bufio"
)

var State = make(chan string)

func handlerEventFull(evento string){

	fmt.Println("FULL-Sensor detected: " ,evento)
	//CALL TO NOTIFICATOR MAYBE

}

func handlerEventPartial(evento string){

	fmt.Println("PARTIAL-Sensor detected: " ,evento)
	//CALL TO NOTIFICATOR MAYBE

}


func ListenEvents(typeofalarm string){
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
		if typeofalarm == cfg.STATE_PART{
			handlerEventPartial(string(line))
		}else{
			handlerEventFull(string(line))
		}

		select {
		case <-State:
			cmd.Process.Kill()
			fmt.Println("saliendodooooo")
			return
		default:
			continue
		}
	}

}

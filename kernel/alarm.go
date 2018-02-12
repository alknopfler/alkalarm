package kernel

import (
	cfg "github.com/alknopfler/alkalarm/config"
	"fmt"
	"strings"
	"os/exec"
	"bufio"
	"github.com/alknopfler/alkalarm/sensors"
	"github.com/alknopfler/alkalarm/control"
	"github.com/alknopfler/alkalarm/states"
)

var State = make(chan string)

func handlerEvent(evento string){

	if sensors.Exists(evento) && states.Query() != cfg.STATE_INAC{
		if sensors.IsPartial(evento) && states.Query() == cfg.STATE_PART{
			Notificate(evento)
			return
		}
		if states.Query() == cfg.STATE_FULL{
			Notificate(evento)
			return
		}
	}else if control.Exists(evento){
		if control.QueryTypeOf(evento) == cfg.STATE_INAC && states.Query() != cfg.STATE_INAC{
			states.Update(cfg.STATE_INAC)
			State <- "stop"
			return
		}
		if control.QueryTypeOf(evento) == cfg.STATE_FULL && states.Query() != cfg.STATE_FULL{
			if states.Query()== cfg.STATE_INAC{
				states.Update(cfg.STATE_FULL)
				go ListenEvents()
				return
			}else if states.Query()== cfg.STATE_PART{
				State <- "stop"  //primero paro y despues lanzo con full
				states.Update(cfg.STATE_FULL)
				go ListenEvents()
				return
			}
		}
		if control.QueryTypeOf(evento) == cfg.STATE_PART && states.Query() != cfg.STATE_PART {
			if states.Query() == cfg.STATE_INAC{
				states.Update(cfg.STATE_PART)
				go ListenEvents()
				return
			}else if states.Query()== cfg.STATE_FULL{
				State <- "stop"  //primero paro y despues lanzo con part
				states.Update(cfg.STATE_PART)
				go ListenEvents()
				return
			}
		}
	}
	//si no es sensor ni control, o la alarma esta inactiva, le dejo pasar sin notificar ni hacer nada
}

func ListenEvents(){
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

		select {
		case <-State:
			cmd.Process.Kill()
			return
		default:
			continue
		}
	}

}

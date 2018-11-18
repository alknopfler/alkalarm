package kernel

import (
	cfg "github.com/alknopfler/alkalarm/config"
	"log"
	"strings"
	"os/exec"
	"bufio"
	"github.com/alknopfler/alkalarm/sensors"
	"github.com/alknopfler/alkalarm/control"
	"github.com/alknopfler/alkalarm/states"
	"time"
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
			return
		}
		if control.QueryTypeOf(evento) == cfg.STATE_FULL && states.Query() != cfg.STATE_FULL{
			if states.Query()== cfg.STATE_INAC{
				time.Sleep(10 * time.Second)
				states.Update(cfg.STATE_FULL)
				return
			}else if states.Query()== cfg.STATE_PART{
				states.Update(cfg.STATE_FULL)
				return
			}
		}
		if control.QueryTypeOf(evento) == cfg.STATE_PART && states.Query() != cfg.STATE_PART {
			if states.Query() == cfg.STATE_INAC{
				time.Sleep(10 * time.Second)
				states.Update(cfg.STATE_PART)
				return
			}else if states.Query()== cfg.STATE_FULL{
				states.Update(cfg.STATE_PART)
				return
			}
		}
	}
	//si no es sensor ni control, o la alarma esta inactiva, le dejo pasar sin notificar ni hacer nada
}

func ListenEvents(){
	cmdName := "python -u " + cfg.PROJECT_PATH + cfg.PYGPIO
	cmdArgs := strings.Fields(cmdName)
	cmd := exec.Command(cmdArgs[0], cmdArgs[1:]...)
	stdout, _ := cmd.StdoutPipe()
	cmd.Start()

	oneByte := make([]byte,0)
	for {
		_, err := stdout.Read(oneByte)
		if err != nil {
			log.Printf(err.Error())
			break
		}
		r := bufio.NewReader(stdout)
		line,_, _ := r.ReadLine()

		handlerEvent(string(line))
	}

	cmd.Process.Wait()

}

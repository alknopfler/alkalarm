package kernel

import (
	cfg "github.com/alknopfler/alkalarm/config"
	"github.com/alknopfler/alkalarm/database"
	"fmt"
	"strings"
	"os/exec"
	"bufio"
)

var State = make(chan string)

func GetGlobalState()  string{
	var state string
	db,err := database.InitDB()
	if err != nil {
		fmt.Println("Error initiating DB in Query Control")
		return state
	}
	defer db.Close()
	rows, err := db.Query(cfg.GLOBAL_STATE_QUERY)
	if err != nil { return state}
	defer rows.Close()

	if rows.Next() {
		err2 := rows.Scan(&state)
		if err2 != nil { return state }
	}
	return state
}

func UpdateGlobalState(newState string)  error{
	db,err := database.InitDB()
	if err != nil {
		fmt.Println("Error initiating DB in Query Control")
		return err
	}
	defer db.Close()
	err=database.Operate(db,cfg.GLOBAL_STATE_UPDATE,newState)
	if err!=nil{
		return err
	}
	return nil
}


func handlerEvent(evento string){

	fmt.Println("Sensor detected: " ,evento)
	//CALL TO NOTIFICATOR MAYBE

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
		if "stop" == <-State{
			break
		}
	}
	cmd.Process.Kill()
}

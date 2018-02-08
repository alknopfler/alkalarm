package main

import (
"bufio"
"fmt"
"os/exec"
"strings"
"github.com/alknopfler/alkalarm/config"
"github.com/alknopfler/alkalarm/database"
)


func handlerEvent(evento string){
	if evento == "3462412"{
		fmt.Println("ha pulsado cerrar")
	}
}

func listenEvents(){
	cmdName := "python -u" + config.PROJECT_PATH + config.PYGPIO
	cmdArgs := strings.Fields(cmdName)

	cmd := exec.Command(cmdArgs[0], cmdArgs[1:len(cmdArgs)]...)
	stdout, _ := cmd.StdoutPipe()
	cmd.Start()
	oneByte := make([]byte,0)
	//TODO convertir en daemon y el python tambien
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

func init() {
	database.CreateSchemas()
}

func main() {

	fmt.Println("main")

}

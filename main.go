package main

import (
"bufio"
"fmt"
"os/exec"
"strings"
)
func handlerEvent(evento string){
	if evento == "3462412"{
		fmt.Println("ha pulsado cerrar")
	}
}

func main() {
	cmdName := "python -u /opt/alkalarm/_433.py"
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

package serverStoper

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
)

func init() {
	if runtime.GOOS != "linux" {
		println("The program only support Linux servers")
		os.Exit(1)
	}
}

func StopServer(screen_id int, time_before_stop int) {
	if screen_id != 0 {
		stoperScriptCreator(screen_id, time_before_stop)
	}

	_, err := os.Open("stoper")
	if err == nil {
		cmd := exec.Command("/bin/sh", "stoper")
		err := cmd.Run()
		if err != nil {
			log.Fatalf("cmd.Run(stoper) failed with %s\n", err)
		}
	} else {
		log.Fatalf("stoper script can't be found and you didn't get parameters to create new one!")
	}
}

func stoperScriptCreator(screen_id int, time_before_stop int) {
	stoper_script, err := os.Create("stoper")
	if err != nil {
		log.Fatalf("File can't be created; %s\n", err)
	}

	defer stoper_script.Close()

	_, err = stoper_script.WriteString(fmt.Sprintf("#!/bin/bash\nscreen -S %v -p 0 -X stuff 'say Server will fall after %v sec^M'\nsleep %v\nscreen -S %v -p 0 -X stuff 'stop^M'", screen_id, time_before_stop, time_before_stop, screen_id))
	if err != nil {
		log.Fatalf("File can't be writed; %s\n", err)
	}

	stoper_script.Sync()
}

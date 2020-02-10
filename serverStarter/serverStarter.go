package serverStarter

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

func StartServer(screen_id int) {
	if screen_id != 0 {
		starterScriptCreator(screen_id)
	}

	_, err := os.Open("starter")
	if err == nil {
		cmd := exec.Command("/bin/sh", "starter")
		err := cmd.Run()
		if err != nil {
			log.Fatalf("cmd.Run(starter) failed with %s\n", err)
		}
	} else {
		log.Fatalf("starter script can't be found and you didn't get parameters to create new one!")
	}
}

func starterScriptCreator(screen_id int) {
	starter_script, err := os.Create("starter")
	if err != nil {
		log.Fatalf("File can't be created; %s\n", err)
	}

	defer starter_script.Close()

	_, err = starter_script.WriteString(fmt.Sprintf("#!/bin/bash\nscreen -S %v -p 0 -X stuff './start^M'", screen_id))
	if err != nil {
		log.Fatalf("File can't be writed; %s\n", err)
	}

	starter_script.Sync()
}

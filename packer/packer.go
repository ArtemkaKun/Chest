package packer

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"
)

func MakePack(files_path string) {
	if files_path != "" {
		packerScriptCreator(files_path)
	}

	_, err := os.Open("packer")
	if err == nil {
		cmd := exec.Command("/bin/sh", "packer")
		err := cmd.Run()
		if err != nil {
			log.Fatalf("cmd.Run(packer) failed with %s\n", err)
		}
	} else {
		log.Fatalf("packer script can't be found and you didn't get parameters to create new one!")
	}
}

func packerScriptCreator(files_path string) {
	packer_script, err := os.Create("packer")
	if err != nil {
		log.Fatalf("File can't be created; %s\n", err)
	}

	defer packer_script.Close()

	dt := time.Now()
	_, month, day := dt.Date()
	_, err = packer_script.WriteString(fmt.Sprintf("#!/bin/bash\ntar -cf %v_%v.tar %v\n7z a %v_%v.7z %v_%v.tar", month.String(), day, files_path, month.String(), day, month.String(), day))
	if err != nil {
		log.Fatalf("File can't be writed; %s\n", err)
	}

	packer_script.Sync()
}

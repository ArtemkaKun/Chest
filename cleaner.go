package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func CleanBackups(file_name string) {
	_, err := os.Open("cleaner")
	if err == nil {
		runScript()
	} else {
		cleanerScriptCreator(file_name)
		runScript()
	}
}

func runScript() {
	cmd := exec.Command("/bin/sh", "cleaner")
	err := cmd.Run()
	if err != nil {
		log.Fatalf("cmd.Run(cleaner) failed with %s\n", err)
	}
}

func cleanerScriptCreator(file_name string) {
	cleaner_script, err := os.Create("cleaner")
	if err != nil {
		log.Fatalf("File can't be created; %s\n", err)
	}

	defer cleaner_script.Close()

	_, err = cleaner_script.WriteString(fmt.Sprintf("#!/bin/bash\n rm -f *.tar *.7z"))
	if err != nil {
		log.Fatalf("File can't be writed; %s\n", err)
	}

	cleaner_script.Sync()
}

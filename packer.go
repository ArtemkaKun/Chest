package main

import (
	"fmt"
	"github.com/metakeule/fmtdate"
	"log"
	"os"
	"os/exec"
	"time"
)

func MakePack(files_path string, rar bool) (pack_name string) {
	if files_path != "" {
		packerScriptCreator(files_path, rar)
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

	dt := time.Now()

	//fmtdate package make normal date form that is tha same with bash date form
	if !rar {
		pack_name = fmt.Sprintf("%v.7z", fmtdate.FormatDate(dt))
	} else {
		pack_name = fmt.Sprintf("%v.rar", fmtdate.FormatDate(dt))
	}

	return pack_name
}

func packerScriptCreator(files_path string, rar bool) {
	packer_script, err := os.Create("packer")
	if err != nil {
		log.Fatalf("File can't be created; %s\n", err)
	}

	defer packer_script.Close()
	//pack to .tar archive and than to .7z or .rar archive. Why? Go to Google and read about .tar format
	if !rar {
		_, err = packer_script.WriteString(fmt.Sprintf("#!/bin/bash\nyear=$(date +'%%Y')\nmonth=$(date +'%%m')\nday=$(date +'%%d')\ntar -cf $year-$month-$day.tar %v\n7z a $year-$month-$day.7z $year-$month-$day.tar", files_path))
		if err != nil {
			log.Fatalf("File can't be writed; %s\n", err)
		}
	} else {
		_, err = packer_script.WriteString(fmt.Sprintf("#!/bin/bash\nyear=$(date +'%%Y')\nmonth=$(date +'%%m')\nday=$(date +'%%d')\nrar a -r $year-$month-$day.rar %v", files_path))
		if err != nil {
			log.Fatalf("File can't be writed; %s\n", err)
		}
	}

	packer_script.Sync()
}

package main

import (
	"flag"
	"fmt"
	"sync"
	"time"
)

func main() {
	//var that contains the id of screen session
	new_screen_id := flag.Int("ns", 0, "Make new server-stoper and server-starter scripts (place here screen ID)")

	//var that contains delay time before a Minecraft server will be stop in seconds
	new_stoper_time := flag.Int("d", 30, "Make new server-stoper script (place here time before stop)")

	//var that contains a max number of files that can contain a backup folder
	max_files := flag.Int("max", 3, "Set the max number of files in backup folder")

	//var that contains a path to folder you want to backup (Minecraft server folder)
	new_server_path := flag.String("f", "", "Make new packer script (place here path to your Minecraft server files)")

	//var that contains a path(or ID) to folder (on Google Drive for example) where program will be upload backups
	backup_folder_path := flag.String("bf", "", "Path to backup folder (place here id of folder if you use Google Drive)")

	//if true, program do backup process only one time and you don't need to set hours and minutes where to make a backup
	test_mode := flag.Bool("test", false, "Turn on test mode, that allow not wait for backup time")
	new_hour := flag.Int("h", 00, "Hour when program need to backup")
	new_minute := flag.Int("m", 00, "Minute when program need to backup")
	flag.Parse()

	backup_time := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), *new_hour, *new_minute, 0, 0, time.Now().UTC().Location())
	if !*test_mode {
		//need endless circle here because program wait for backup time and check time every minute
		for true {
			if (time.Now().Hour() == backup_time.Hour()) && (time.Now().Minute() == backup_time.Minute()) {
				backupProcess(*new_screen_id, *new_stoper_time, *new_server_path, *backup_folder_path, *max_files)
				time.Sleep(time.Minute)
				continue
			} else {
				continue
			}
		}
	} else {
		backupProcess(*new_screen_id, *new_stoper_time, *new_server_path, *backup_folder_path, *max_files)
	}
}

func backupProcess(new_screen_id int, new_stoper_time int, new_server_path string, backup_folder_path string, max_files int) {
	fmt.Println("The server will be stopped after delay")
	StopServer(new_screen_id, new_stoper_time)

	fmt.Println("The server was stopped. Process of backup making was started...")
	new_pack := MakePack(new_server_path)

	fmt.Println("Backup file was successfully created! Process of uploading was started...")
	var upload_goroutine sync.WaitGroup
	upload_goroutine.Add(1)
	go UploadPack(new_pack, &upload_goroutine, backup_folder_path)

	fmt.Println("Starting the server...")
	StartServer(new_screen_id)
	fmt.Println("The server started")

	//need to wait while uploading backup will be completed
	upload_goroutine.Wait()
	
	CheckOld(backup_folder_path, max_files)
	CleanBackups(new_pack)
}

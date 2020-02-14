package main

import (
	"./deleter"
	"./packer"
	starter "./serverStarter"
	stoper "./serverStoper"
	"./uploader"
	"flag"
	"fmt"
	"sync"
	"time"
)

func main() {
	new_screen_id := flag.Int("ns", 0, "Make new server-stoper and server-starter scripts (place here screen ID)")
	new_stoper_time := flag.Int("d", 30, "Make new server-stoper script (place here time before stop)")
	max_files := flag.Int("max", 3, "Set the max number of files in backup folder")
	new_server_path := flag.String("f", "", "Make new packer script (place here path to your Minecraft server files)")
	backup_folder_path := flag.String("bf", "", "Path to backup folder (place here id of folder if you use Google Drive)")
	test_mode := flag.Bool("test", false, "Turn on test mode, that allow not wait for backup time")
	new_hour := flag.Int("h", 00, "Hour when program need to backup")
	new_minute := flag.Int("m", 00, "Minute when program need to backup")
	flag.Parse()

	backup_time := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), *new_hour, *new_minute, 0, 0, time.Now().UTC().Location())
	if !*test_mode {
		for true {
			if (time.Now().Hour() == backup_time.Hour()) && (time.Now().Minute() == backup_time.Minute()) {
				fmt.Println("The server will be stopped after delay")
				stoper.StopServer(*new_screen_id, *new_stoper_time)

				fmt.Println("The server was stopped. Process of backup making was started...")
				new_pack := packer.MakePack(*new_server_path)

				fmt.Println("Backup file was successfully created! Process of uploading was started...")
				var upload_goroutine sync.WaitGroup
				upload_goroutine.Add(1)
				go uploader.UploadPack(new_pack, &upload_goroutine, *backup_folder_path)

				fmt.Println("Starting the server...")
				starter.StartServer(*new_screen_id)
				fmt.Println("The server started")

				upload_goroutine.Wait()

				deleter.CheckOld(*backup_folder_path, *max_files)
				continue
			} else {
				continue
			}
		}
	} else {
		fmt.Println("The server will be stopped after delay")
		stoper.StopServer(*new_screen_id, *new_stoper_time)

		fmt.Println("The server was stopped. Process of backup making was started...")
		new_pack := packer.MakePack(*new_server_path)

		fmt.Println("Backup file was successfully created! Process of uploading was started...")
		var upload_goroutine sync.WaitGroup
		upload_goroutine.Add(1)
		go uploader.UploadPack(new_pack, &upload_goroutine, *backup_folder_path)

		fmt.Println("Starting the server...")
		starter.StartServer(*new_screen_id)
		fmt.Println("The server started")

		upload_goroutine.Wait()

		deleter.CheckOld(*backup_folder_path, *max_files)
	}
}

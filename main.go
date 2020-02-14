package main

import (
	"./deleter"
	//"./packer"
	//starter "./serverStarter"
	//stoper "./serverStoper"
	//"./uploader"
	"flag"
	//"fmt"
	//"sync"
)

func main() {
	//new_screen_id := flag.Int("ns", 0, "Make new server-stoper and server-starter scripts (place here screen ID)")
	//new_stoper_time := flag.Int("d", 30, "Make new server-stoper script (place here time before stop)")
	//new_server_path := flag.String("f", "", "Make new packer script (place here path to your Minecraft server files)")
	backup_folder_path := flag.String("bf", "", "Path to backup folder (place here id of folder if you use Google Drive)")
	max_files := flag.Int("max", 3, "Set the max number of files in backup folder")
	flag.Parse()

	//fmt.Println("The server will be stopped after delay")
	//stoper.StopServer(*new_screen_id, *new_stoper_time)
	//
	//fmt.Println("The server was stopped. Process of backup making was started...")
	//new_pack := packer.MakePack(*new_server_path)
	//
	//fmt.Println("Backup file was successfully created! Process of uploading was started...")
	//var upload_goroutine sync.WaitGroup
	//upload_goroutine.Add(1)
	//go uploader.UploadPack(new_pack, &upload_goroutine, *backup_folder_path)
	//
	//fmt.Println("Starting the server...")
	//starter.StartServer(*new_screen_id)
	//fmt.Println("The server started")
	//
	//upload_goroutine.Wait()

	deleter.CheckOld(*backup_folder_path, *max_files)
}

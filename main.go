package main

import (
	"./packer"
	starter "./serverStarter"
	stoper "./serverStoper"
	"./uploader"
	"flag"
)

func main() {
	new_screen_id := flag.Int("nssid", 0, "Make new server-stoper and server-starter scripts (place here screen ID)")
	new_stoper_time := flag.Int("nstime", 30, "Make new server-stoper script (place here time before stop)")
	new_server_path := flag.String("sf", "", "Make new packer script (place here path to your Minecraft server files)")
	flag.Parse()

	stoper.StopServer(*new_screen_id, *new_stoper_time)
	uploader.UploadPack(packer.MakePack(*new_server_path))
	starter.StartServer(*new_screen_id)
}

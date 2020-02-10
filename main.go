package main

import (
	starter "./serverStarter"
	stoper "./serverStoper"
	"flag"
)

func main() {
	new_screen_id := flag.Int("nssid", 0, "Make new server-stoper and server-starter scripts (place here screen ID)")
	new_stoper_time := flag.Int("nstime", 30, "Make new server-stoper script (place here time before stop)")
	flag.Parse()

	stoper.StopServer(*new_screen_id, *new_stoper_time)
	starter.StartServer(*new_screen_id)
}

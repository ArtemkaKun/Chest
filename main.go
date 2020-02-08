package main

import (
	stoper "./serverStoper"
)

const SCREEN_ID string = "14188"

func main() {
	stoper.StopServer(SCREEN_ID, 30)
}

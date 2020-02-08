package main

import (
	stoper "./serverStoper"
)

const SCREEN_ID string = "6154"

func main() {
	stoper.StopServer(SCREEN_ID)
}

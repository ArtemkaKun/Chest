package serverStoper

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
)

func init() {
	if runtime.GOOS != "linux" {
		println("The program only support Linux servers")
		os.Exit(1)
	}
}

func StopServer(screen_id string) {
	cmd := exec.Command("/bin/sh", "stoper")
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
		log.Fatalf("cmd.Run(screen -r) failed with %s\n", err)
	}

}

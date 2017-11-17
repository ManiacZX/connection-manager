package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func main() {
	started := time.Now()

	wvdial()

	for {
		if time.Since(started) > 5*time.Minute {
			break
		}
		time.Sleep(30 * time.Second)
	}

	reboot()
}

func wvdial() {
	cmd := exec.Command("wvdial")
	err := cmd.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error executing wvdial: %v", err)
	}
}

func reboot() {
	cmd := exec.Command("reboot")
	err := cmd.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error rebooting: %v", err)
		os.Exit(1)
	}
}

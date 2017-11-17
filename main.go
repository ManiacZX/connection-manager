package main

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"time"
)

func main() {
	started := time.Now()

	terminate := make(chan string)

	go wvdial(terminate)
	go netmon(terminate)

	<-terminate

	for {
		if time.Since(started) > 5*time.Minute {
			break
		}
		time.Sleep(30 * time.Second)
	}

	reboot()
}

func wvdial(terminate chan<- string) {
	cmd := exec.Command("wvdial")
	err := cmd.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error executing wvdial: %v", err)
	}
	terminate <- "wvdial"
}

func netmon(terminate chan<- string) {
	errCount := 0
	for {
		_, err := http.Get("https://www.google.com")
		if err != nil {
			errCount++
			if errCount > 5 {
				terminate <- "netmon"
				break
			}
		} else {
			errCount = 0
		}
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

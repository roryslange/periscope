package main

import (
	"log"
	"os/exec"
	"time"

	"github.com/roryslange/periscope/data-collection"
	"github.com/roryslange/periscope/process-control"
	"golang.org/x/sys/unix"
)

func main() {
	var command = "stress"
	var args = []string{"--cpu", "4",
		"--io", "2",
		"--vm", "1",
		"--vm-bytes", "64M",
		"--timeout", "5s"}

	target := exec.Command(command, args...)
	stdout, _ := target.StdoutPipe()
	stderr, _ := target.StderrPipe()
	
	log.Println("start")
	err := target.Start()

	if (err != nil) {
		log.Fatal(err)
	}

	// todo: combine these prints, maybe just organize it better
	go processcontrol.PrintCmdReaderOutput(&stdout)
	go processcontrol.PrintCmdReaderOutput(&stderr)

	go datacollection.PrintDetails(unix.RUSAGE_SELF, time.Millisecond)


	//wait for it to finish
	target.Wait()
	log.Println("done")

	//do diagnostics summary
}

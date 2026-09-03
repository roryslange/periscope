package main

import (
	"log"
	"os/exec"

	"github.com/roryslange/periscope/process-control"
	"github.com/roryslange/periscope/data-collection"
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

	go processcontrol.PrintCmdReaderOutput(&stdout)
	go processcontrol.PrintCmdReaderOutput(&stderr)
	go datacollection.PrintCpuTime()


	//wait for it to finish
	target.Wait()
	log.Println("done")

	//do diagnostics summary
}

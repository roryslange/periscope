package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os/exec"
	"time"

	"golang.org/x/sys/unix"
)

func main() {
	cmd := exec.Command("ls", "-al")

	stdout, _ := cmd.StdoutPipe()
	err := cmd.Start()

	if (err != nil) {
		log.Fatal(err)
	}

	go printCmdOutput(stdout)
	go getSomeDiagnostics()


	//wait for it to finish
	cmd.Wait()

	//do diagnostics summary
}

func getSomeDiagnostics() {
	for {
		var rusage unix.Rusage
		unix.Getrusage(unix.RUSAGE_SELF, &rusage) //i think self is the wrong process here

		fmt.Printf("User CPU: %v\tSystem CPU: %v\n", rusage.Utime, rusage.Stime)
		time.Sleep(time.Millisecond)
	}
}

func printCmdOutput(stdout io.Reader) {
	scanner := bufio.NewScanner(stdout)

		err := scanner.Err()
		if (err != nil) {
			log.Fatal(err)
		}

		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
}
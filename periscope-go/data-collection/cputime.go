package datacollection

import (
	"log"
	"time"

	"golang.org/x/sys/unix"
)

func PrintDetails(pid int, refreshRateNs time.Duration) {
	for {
		rusage := getCpuTime(pid)

		log.Printf("\nUser CPU: %+v\tSystem CPU: %+v\n", rusage.Utime, rusage.Stime)
		time.Sleep(refreshRateNs)
	}
}

func getCpuTime(pid int) unix.Rusage {
	var rusage unix.Rusage
	unix.Getrusage(pid, &rusage)
	return rusage
}
package datacollection

import (
	"log"
	"time"

	"golang.org/x/sys/unix"
)

func PrintCpuTime() {
	for {
		var rusage unix.Rusage
		unix.Getrusage(unix.RUSAGE_SELF, &rusage) //i think self is the wrong process here

		log.Printf("\nUser CPU: %+v\tSystem CPU: %+v\n\n", rusage.Utime, rusage.Stime)
		time.Sleep(time.Millisecond)
	}
}
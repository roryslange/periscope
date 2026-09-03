package processcontrol

import (
	"bufio"
	"io"
	"log"
)

func PrintCmdReaderOutput(outputPipe *io.ReadCloser) {
	scanner := bufio.NewScanner(*outputPipe)
	for scanner.Scan() {
		log.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
}
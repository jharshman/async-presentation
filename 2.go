package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

	/*
	This is better but not great. The goroutine still doesn't close gracefully.
	The program exits because the execution falls out the end of main.
	There is no real communication between the goroutine and the main process.
	Let's fix that.
	 */
// START OMIT
func main() {
	closeChan := make(chan os.Signal, 1)
	signal.Notify(closeChan, syscall.SIGINT) // HLchannel
	go boring()
	<-closeChan // HLchannel
}
// END OMIT

func boring() {
	for {
		fmt.Println("hello gophers!")
		<-time.After(time.Second * 2)
	}
}


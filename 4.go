package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)
	/*
	Small changes here...
	We now return two channels from the boring() function.
	And once we signal the goroutine to close, main waits for an acknoledgement
	from the goroutine that it indeed has finished execution.

	We have now successfully gotten rid of any sleeping in our code. But more importantly,
	we have intelligent communication between our main process and the single goroutine it spawns.
	*/

func main() {
	closeChan := make(chan os.Signal, 1)
	signal.Notify(closeChan, syscall.SIGINT)

	sig, ack := boring()
	<-closeChan

	sig<-1
	<-ack // HLack
}

func boring() (chan int, chan int) { // HLack
	sig := make(chan int, 1)
	ack := make(chan int, 1) // HLack
	go func() {
	LOOP:
		for {
			select {
			case <-sig:
				break LOOP
			case <-time.After(time.Second * 2):
				fmt.Println("hello gophers!")
			}
		}
		fmt.Println("goodbye")
		ack<-1 // HLack
	}()
	return sig, ack // HLack
}


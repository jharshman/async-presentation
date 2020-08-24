package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)
	/*
	Now there are some significant changes here...
	The first being that the boring() function is no longer being called in main
	as `go boring()`. Instead, it is a function that then runs a goroutine, passing back
	a channel, to communicate with that goroutine.

	We still listen for SIGINT, but instead of allowing the program to close, we send a value
	to the channel that was returned to us from the boring() function. This value is then received
	by the goroutine, and instructs it to finish execution.

	We still have to hold main() open until the fmt.Println("goodbye") in the goroutine can be called.
	Perhaps we can improve upon this aspect as well.

	And what if errors occur in the goroutine? How can we be informed of them?
	*/

func main() {
	closeChan := make(chan os.Signal, 1)
	signal.Notify(closeChan, syscall.SIGINT)

	sig := boring() // HLsig

	<-closeChan
	sig<-1 // HLsig

  // Hold main from exiting for 2 miliseconds
  // so boring() can print "goodbye".
	<-time.After(time.Millisecond*2)
}

func boring() chan int {
	sig := make(chan int, 1)
	go func() {
LOOP:
		for {
			select { // HLselect
			case <-sig: // HLselect
				break LOOP // HLselect
			case <-time.After(time.Second * 2): // HLselect
				fmt.Println("hello gophers!") // HLselect
			} // HLselect
		}
		fmt.Println("goodbye")
	}()
	return sig // HLselect
}


package main

import (
	"fmt"
	"time"
)
	// run some boring task in a goroutine

	/*
	without this timeout, the goroutine launched above wouldn't even get a chance to print to
	the console before program execution terminated. How do we make this program better?
	A good thought is to listen and wait for specific signals like SIGINT or SIGKILL and then
	allow the program to terminate.
	 */
// START OMIT
func main() {
	go boring() // HLboring
	<-time.After(time.Second * 10)
}

func boring() {
	for {
		fmt.Println("hello gophers!")
		<-time.After(time.Second * 2)
	}
}
// END OMIT


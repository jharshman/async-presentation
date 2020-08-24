package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
  "sync"
	"time"
)
	/*
  You can alternatively use a WaitGroup to wait for the goroutine to exit.
	*/

func main() {
	closeChan := make(chan os.Signal, 1)
	signal.Notify(closeChan, syscall.SIGINT)

  var wg sync.WaitGroup
  wg.Add(1) // HLwg
	sig := boring(&wg)
	<-closeChan

	sig<-1
	wg.Wait() // HLwg
}

func boring(wg *sync.WaitGroup) chan int { // HLwg
	sig := make(chan int, 1)
	go func() {
    defer wg.Done() // HLwg
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
	}()
	return sig
}


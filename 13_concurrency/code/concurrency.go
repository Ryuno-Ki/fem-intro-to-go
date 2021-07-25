package main

import (
	"fmt"
	"sync"
	"time"
)

// Used to manage the order in which threads execute
var wg sync.WaitGroup

// Say something three times
func Say(s string) {
	for i := 0; i < 3; i++ {
		fmt.Println(s)
		time.Sleep(time.Millisecond * 300)
	}
}

func handlePanic () {
  if r := recover(); r != nil {
    fmt.Println("Panic")
  }
}

func printStuff () {
  // Tell waitgroup, this function is completed
  defer wg.Done()
  // Handle failures or otherwise the thread won't be popped from the stack
  // causing the waitgroup to wait forever
  defer handlePanic()

  for i := 0; i < 3; i++ {
    fmt.Println(i)
    time.Sleep(time.Millisecond * 300)
  }
}

func main() {
	// Turn function into a thread
	//go Say("Hello")
	// If you turn this into a thread as well, nothing happens, because it
	// waits for the next line to finish first.
	//Say("There")

	// Initialise with an integer
	wg.Add(1)
	go printStuff()
	wg.Wait()
	// Next step would be to look into channels for communication between
	// threads
}

package main

import (
	"fmt"
	"time"
)

// START OMIT
func main() {
	done := make(chan bool) // HL
	go longRunning(done)
	fmt.Println("hi")
	<-done // HL
	fmt.Println("longRunning done!")
}

func longRunning(done chan bool) {
	time.Sleep(3 * time.Second)
	done <- true // HL
}

// END OMIT

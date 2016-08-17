package main

import (
	"fmt"
	"time"
)

// START OMIT
func main() {
	go longRunning() // HL
	fmt.Println("hi")
}

func longRunning() {
	time.Sleep(3 * time.Second)
}

// END OMIT

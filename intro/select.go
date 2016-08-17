package main

import (
	"fmt"
	"math/rand"
	"time"
)

// START OMIT
func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go calculate(ch1)
	go calculate(ch2)
	for {
		select { // HL
		case v := <-ch1: // HL
			fmt.Printf("From ch1: %d\n", v)
		case v := <-ch2: // HL
			fmt.Printf("From ch2: %d\n", v)
		}
	}
}

func calculate(ch chan int) {
	for {
		time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
		ch <- 42 // HL
	}
}

// END OMIT

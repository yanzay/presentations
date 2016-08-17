package main

import (
	"fmt"
	"time"
)

// START OMIT
func main() {
	res := make(chan int)
	go calculateTheAnswer(res)
	fmt.Println("Doing some job here...")
	result := <-res // HL
	fmt.Printf("The Answer is: %d\n", result)
}

func calculateTheAnswer(res chan int) {
	time.Sleep(3 * time.Second)
	res <- 42 // HL
}

// END OMIT

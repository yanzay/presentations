package main

import "fmt"

// START OMIT
func main() {
	res := make(chan int, 3) // HL
	res <- 1
	res <- 2
	res <- 3
	fmt.Println(<-res, <-res, <-res)
}

// END OMIT

package main

import "fmt"

// START OMIT
type Sayer interface {
	Say() string
}

type User struct{}

func (u User) Say() string {
	return "Hello"
}

func main() {
	var user Sayer
	user = User{}
	fmt.Println(user.Say())
}

// END OMIT

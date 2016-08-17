package main

import "fmt"

// START OMIT
type User struct {
	Name string
	Age  int
}

func (u User) GetAge() int {
	return u.Age
}

func main() {
	user := User{
		Name: "Alexey",
		Age:  18,
	}
	fmt.Println(user.Name)
	fmt.Println(user.GetAge())
}

// END OMIT

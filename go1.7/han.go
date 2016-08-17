package main

import (
	"fmt"
	"reflect"
)

func main() {
	stringType := reflect.TypeOf("abc")
	chanType := reflect.ChanOf(3, stringType)
	fmt.Println(chanType.String())
}

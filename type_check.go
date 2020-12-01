package main

import (
	"fmt"
	"reflect"
)

func main() {

	input := "12+34"

	fmt.Println(reflect.TypeOf(input))
	fmt.Println(reflect.TypeOf(string(input[2])))
	fmt.Println(reflect.TypeOf(input[2:3]))
}

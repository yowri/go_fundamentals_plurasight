package main

import (
	"fmt"
	"reflect"
)

var (
	// declaring types u can do like this

	// name, course string
	// module       float64

	// or like this

	// name, course, module = "nigel", "Docker Deep Dive", 3.2

	// but this is more readable
	name   = "nigel"
	course = "Docker Deep Dive"
	module = 3.2
)

func main() {
	fmt.Println("Name is set to", name, "and is of type ", reflect.TypeOf(name))
	fmt.Println("Name is set to", module, "and is of type ", reflect.TypeOf(module))
}

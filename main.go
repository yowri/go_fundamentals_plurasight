package main

import (
	"fmt"
	"os"
)

// // Package level variables are global
// var (
// 	// declaring types u can do like this
//
// 	// name, course string
// 	// module       float64
//
// 	// or like this
//
// 	// name, course, module = "nigel", "Docker Deep Dive", 3.2
//
// 	// but this is more readable
// 	name   = "nigel"
// 	course = "Docker Deep Dive"
// 	module = 3.2
// )

func main() {
	// // This declaring only works in function.
	// // Variables at the function level must be used.
	name := os.Getenv("USERNAME")
	course := "Docker Deep Dive"
	// module := 3.2

	// // creating a pointer value
	// ptr := &module
	//
	// fmt.Println("Name is set to", name, "and is of type ", reflect.TypeOf(name))
	// fmt.Println("Name is set to", module, "and is of type ", reflect.TypeOf(module))
	//
	// // showing a pointer value
	// fmt.Println("The memory adress of the *module* is", ptr, "and the value of the memory is", *ptr)

	// // Passing by value
	fmt.Println("\nHi", name, "you're currently watching", course)

	changeCourse(&course)

	fmt.Println("\nYou are now watching course", course)
}

func changeCourse(course *string) {
	*course = "First look: Native Docker Clustering"

	fmt.Println("\nTrying to change your course to", *course)
}

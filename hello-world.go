package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Hello form", runtime.GOOS)
}

package main

import (
	"fmt"
)

func ReturnFunc() func() {
	return func() {
		fmt.Println("I'm a function")
	}
}

func main() {
	f := ReturnFunc()
	f()
}
